package ssh
import (

	"time"

	"golang.org/x/crypto/ssh"
)
// ,USER, PASSWORD, DIAL_PORT, BIND_PORT, HOST_ADDR, HOST_PORT,
type Queue_Request struct {
	ID				int `json:"id"`
	USERNAME       string `json:"name"`
	EMAIL      		string `json:"email"`
	LOCATION   		string `json:"location"`
	CREATED_AT 		string `json:"created_at"`
	DIAL_PORT        string `json:"dial_port"`
	BIND_PORT        string `json:"bind_port"`
	HOST_ADDR       string `json:"host_addr"`
	HOST_PORT        string `json:"host_port"`
	PASSWORD        string `json:"pwd"`
}
// https://github.com/dsnet/sshtunnel

const (
	TIMEOUT_RETRY = 30
	SSH_MODE = 1
	SSH_BAD_REQUEST = "BAD_REQUEST"
	SSH_NUM_TIMOUT = 3
)

var (
	UnFulfilledRequests map[string]int
)
type TunnelConfig struct {
	// LogFile is where the proxy daemon will direct its output log.
	// If the path is empty, then the server will output to os.Stderr.
	LogFile string `json:",omitempty"`

	// KeyFiles is a list of SSH private key files.
	KeyFiles []string

	// KnownHostFiles is a list of key database files for host public keys
	// in the OpenSSH known_hosts file format.
	KnownHostFiles []string

	// KeepAlive sets the keep alive settings for each SSH connection.
	// It is recommended that these values match the AliveInterval and
	// AliveCountMax parameters on the remote OpenSSH server.
	// If unset, then the default is an interval of 30s with 2 max counts.
	KeepAlive *KeepAliveConfig `json:",omitempty"`

	// Tunnels is a list of tunnels to establish.
	// The same set of SSH keys will be used to authenticate the
	// SSH connection for each server.
	Tunnels []struct {
		// Tunnel is a pair of host:port endpoints that can be configured
		// to either operate as a forward tunnel or a reverse tunnel.
		//
		// The syntax of a forward tunnel is:
		//	"bind_address:port -> dial_address:port"
		//
		// A forward tunnel opens a listening TCP socket on the
		// local side (at bind_address:port) and proxies all traffic to a
		// socket on the remote side (at dial_address:port).
		//
		// The syntax of a reverse tunnel is:
		//	"dial_address:port <- bind_address:port"
		//
		// A reverse tunnel opens a listening TCP socket on the
		// remote side (at bind_address:port) and proxies all traffic to a
		// socket on the local side (at dial_address:port).
		Tunnel string

		// Server is a remote SSH host. It has the following syntax:
		//	"user@host:port"
		//
		// If the user is missing, then it defaults to the current process user.
		// If the port is missing, then it defaults to 22.
		Server string

		// KeepAlive is a tunnel-specific setting of the global KeepAlive.
		// If unspecified, it uses the global KeepAlive settings.
		KeepAlive *KeepAliveConfig `json:",omitempty"`
	}
}

type KeepAliveConfig struct {
	// Interval is the amount of time in seconds to wait before the
	// tunnel client will send a keep-alive message to ensure some minimum
	// traffic on the SSH connection.
	Interval uint

	// CountMax is the maximum number of consecutive failed responses to
	// keep-alive messages the client is willing to tolerate before considering
	// the SSH connection as dead.
	CountMax uint
}

type tunnel struct {
	auth          []ssh.AuthMethod
	hostKeys      ssh.HostKeyCallback
	mode          byte // '>' for forward, '<' for reverse
	user          string
	hostAddr      string
	bindAddr      string
	dialAddr      string
	retryInterval time.Duration
	keepAlive     KeepAliveConfig
	//log logger
}