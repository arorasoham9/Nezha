package ssh

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net"
	"os"
	"os/signal"
	"path"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh"
)

func (t tunnel) String() string {
	var left, right string
	mode := "<?>"
	switch t.mode {
	case '>':
		left, mode, right = t.bindAddr, "->", t.dialAddr
	case '<':
		left, mode, right = t.dialAddr, "<-", t.bindAddr
	}
	return fmt.Sprintf("%s@%s | %s %s %s", t.user, t.hostAddr, left, mode, right)
}

func (t tunnel) bindTunnel(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		var once sync.Once // Only print errors once per session
		func() {
			// Connect to the server host via SSH.
			cl, err := ssh.Dial("tcp", t.hostAddr, &ssh.ClientConfig{
				User:            t.user,
				Auth:            t.auth,
				HostKeyCallback: t.hostKeys,
				Timeout:         5 * time.Second,
			})
			if err != nil {
				once.Do(func() { fmt.Printf("(%v) SSH dial error: %v\n", t, err) })
				return
			}
			wg.Add(1)
			go t.keepAliveMonitor(&once, wg, cl)
			defer cl.Close()

			// Attempt to bind to the inbound socket.
			var ln net.Listener
			switch t.mode {
			case '>':
				ln, err = net.Listen("tcp", t.bindAddr)
			case '<':
				ln, err = cl.Listen("tcp", t.bindAddr)
			}
			if err != nil {
				once.Do(func() { fmt.Printf("(%v) bind error: %v\n", t, err) })
				return
			}

			// The socket is binded. Make sure we close it eventually.
			bindCtx, cancel := context.WithCancel(ctx)
			defer cancel()
			go func() {
				cl.Wait()
				cancel()
			}()
			go func() {
				<-bindCtx.Done()
				once.Do(func() {}) // Suppress future errors
				ln.Close()
			}()

			fmt.Printf("(%v) binded tunnel\n", t)
			defer fmt.Printf("(%v) collapsed tunnel\n", t)

			// Accept all incoming connections.
			for {
				cn1, err := ln.Accept()
				if err != nil {
					once.Do(func() { fmt.Printf("(%v) accept error: %v\n", t, err) })
					return
				}
				wg.Add(1)
				go t.dialTunnel(bindCtx, wg, cl, cn1)
			}
		}()

		select {
		case <-ctx.Done():
			return
		case <-time.After(t.retryInterval):
			fmt.Printf("(%v) retrying...\n", t)
		}
	}
}

func (t tunnel) dialTunnel(ctx context.Context, wg *sync.WaitGroup, client *ssh.Client, cn1 net.Conn) {
	defer wg.Done()

	// The inbound connection is established. Make sure we close it eventually.
	connCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-connCtx.Done()
		cn1.Close()
	}()

	// Establish the outbound connection.
	var cn2 net.Conn
	var err error
	switch t.mode {
	case '>':
		cn2, err = client.Dial("tcp", t.dialAddr)
	case '<':
		cn2, err = net.Dial("tcp", t.dialAddr)
	}
	if err != nil {
		fmt.Printf("(%v) dial error: %v", t, err)
		return
	}

	go func() {
		<-connCtx.Done()
		cn2.Close()
	}()

	fmt.Printf("(%v) connection established", t)
	defer fmt.Printf("(%v) connection closed", t)

	// Copy bytes from one connection to the other until one side closes.
	var once sync.Once
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		defer cancel()
		if _, err := io.Copy(cn1, cn2); err != nil {
			once.Do(func() { fmt.Printf("(%v) connection error: %v", t, err) })
		}
		once.Do(func() {}) // Suppress future errors
	}()
	go func() {
		defer wg2.Done()
		defer cancel()
		if _, err := io.Copy(cn2, cn1); err != nil {
			once.Do(func() { fmt.Printf("(%v) connection error: %v", t, err) })
		}
		once.Do(func() {}) // Suppress future errors
	}()
	wg2.Wait()
}

// keepAliveMonitor periodically sends messages to invoke a response.
// If the server does not respond after some period of time,
// assume that the underlying net.Conn abruptly died.
func (t tunnel) keepAliveMonitor(once *sync.Once, wg *sync.WaitGroup, client *ssh.Client) {
	defer wg.Done()
	if t.keepAlive.Interval == 0 || t.keepAlive.CountMax == 0 {
		return
	}

	// Detect when the SSH connection is closed.
	wait := make(chan error, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		wait <- client.Wait()
	}()

	// Repeatedly check if the remote server is still alive.
	var aliveCount int32
	ticker := time.NewTicker(time.Duration(t.keepAlive.Interval) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case err := <-wait:
			if err != nil && err != io.EOF {
				once.Do(func() { fmt.Printf("(%v) SSH error: %v", t, err) })
			}
			return
		case <-ticker.C:
			if n := atomic.AddInt32(&aliveCount, 1); n > int32(t.keepAlive.CountMax) {
				once.Do(func() { fmt.Printf("(%v) SSH keep-alive termination", t) })
				client.Close()
				return
			}
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _, err := client.SendRequest("keepalive@openssh.com", true, nil)
			if err == nil {
				atomic.StoreInt32(&aliveCount, 0)
			}
		}()
	}
}

func loadConfig(MODE int,USERNAME, PASSWORD, DIAL_PORT, BIND_PORT, HOST_ADDR, HOST_PORT string, TIMEOUT_RETRY int) (tunns []tunnel, closer func() error) {
	var auth []ssh.AuthMethod

	auth = append(auth, ssh.Password(PASSWORD))

	var tunn2 tunnel
	tunn2.auth = auth
	tunn2.hostKeys = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	switch MODE{
	case 1:
		tunn2.mode = '>' // '>' for forward, '<' for reverse
	case 2: 
		tunn2.mode = '<' // '>' for forward, '<' for reverse	
	}
	tunn2.user = USERNAME
	tunn2.hostAddr = net.JoinHostPort(HOST_ADDR, HOST_PORT)
	tunn2.bindAddr = "localhost:"+BIND_PORT
	tunn2.dialAddr = "localhost:"+DIAL_PORT
	tunn2.retryInterval = time.Duration(TIMEOUT_RETRY) * time.Second
	//tunn1.keepAlive = *KeepAliveConfig
	tunns = append(tunns, tunn2)

	return tunns, closer
}



// this is not supposed to be  main, I just tried to test it for now in main.go and just pasted it here. Will create the necessary functions soon
func SendOutConnection(rqst []byte, ID string) {
	var request Queue_Request
	err := json.Unmarshal(rqst, &request)
	if err != nil{
		UnFulfilledRequests[ID]++ 
		return
	}
	tunns, closer := loadConfig(SSH_MODE ,request.USERNAME, request.PASSWORD, request.DIAL_PORT, request.BIND_PORT, request.HOST_ADDR, request.HOST_PORT, TIMEOUT_RETRY)
	defer closer()

	// Setup signal handler to initiate shutdown.
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		fmt.Printf("received %v - initiating shutdown\n", <-sigc)
		cancel()
	}()

	// Start a bridge for each tunnel.
	var wg sync.WaitGroup
	fmt.Printf("%s starting\n", path.Base(os.Args[0]))
	defer fmt.Printf("%s shutdown\n", path.Base(os.Args[0]))
	for _, t := range tunns {
		wg.Add(1)
		go t.bindTunnel(ctx, &wg)
	}
	wg.Wait()
}