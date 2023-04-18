package queue

import (
	"fmt"
	"os"

	// "os"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

var queue Queue //singleton implementation

func MakeQueue(api_conn_options, ssh_serv_conn_options *redis.Options) Queue {
	api_connection, err1 := makeQueueConnection(api_conn_options, API_Q_CLI)
	ssh_serv_connection, err2 := makeQueueConnection(ssh_serv_conn_options, SSH_Q_CLI)

	queue = Queue{
		API_CLI:               api_connection,
		SSH_SERV_CLI:          ssh_serv_connection,
		CREATED:               time.Now(),
		ONLINE:                err1 != nil && err2 != nil,
		API_CONN_OPTIONS:      api_conn_options,
		SSH_SERV_CONN_OPTIONS: ssh_serv_conn_options,
	}
	return queue
}

func makeQueueConnection(options *redis.Options, name string) (*redis.Client, error) {
	//begin client
	queueConn := redis.NewClient(options)
	//check if a container exists
	_, err := queueConn.Ping().Result()
	if err != nil {
		fmt.Println("Queue connection:", err)
		os.Exit(2)
	}
	return queueConn, nil
}

func CheckAlive() Queue {
	_, err1 := queue.API_CLI.Ping().Result()
	_, err2 := queue.SSH_SERV_CLI.Ping().Result()
	if err1 != nil || err2 != nil {
		queue.API_CLI.Close()
		queue.SSH_SERV_CLI.Close()
		MakeQueue(queue.API_CONN_OPTIONS, queue.SSH_SERV_CONN_OPTIONS)
	}
	return queue
}

func QueueIsEmpty() bool {
	rslt, err := queue.API_CLI.Keys("*").Result()
	if err == nil {
		return false
	}
	return len(rslt) == 0
}

func shutDownQueue(force bool) int {
	queue.API_CLI.Close()
	queue.SSH_SERV_CLI.Close()
	//run script to restart docker container or smthing TBD, if we are running kubernetes then
	//it might be easier to get this to work
	//check for a successful shutdown
	shutDownCheck := false
	if !shutDownCheck {
		return QUEUE_CONTAINER_SHUTDOWN_UNSUCCESSFUL
	}
	return QUEUE_SHUT_DOWN_SUCCESSFUL
}

func restartQueue(force bool) int {
	if !QueueIsEmpty() && !force {
		return QUEUE_RESTART_FAIL_QUEUE_NOT_EMPTY
	}
	if EmptyQueue() != nil {
		return QUEUE_RESTART_FAIL_COULD_NOT_EMPTY_QUEUE
	}

	if shutDownQueue(force) != QUEUE_SHUT_DOWN_SUCCESSFUL {
		return QUEUE_SHUTDOWN_FAIL
	}
	//run script to restart container
	//check if a container is alive
	MakeQueue(queue.API_CONN_OPTIONS, queue.SSH_SERV_CONN_OPTIONS)

	return QUEUE_RESTART_SUCCESSFUL
}

func EmptyQueue() error {
	return queue.SSH_SERV_CLI.Del("*").Err()
}

func AddRequestToQueue(key string, request Queue_Request) error {
	json_byte_request, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return queue.API_CLI.Set(key, json_byte_request, TTL).Err()
}

func RemoveRequestFromQueue(key string) error {
	return queue.API_CLI.Del(key).Err()
}

func GetRequestFromQueue(key string) (Queue_Request, error) {

	json_byte_request, err := queue.SSH_SERV_CLI.Get(key).Bytes()
	if err != nil {
		return Queue_Request{}, err
	}
	var request Queue_Request
	err = json.Unmarshal([]byte(json_byte_request), &request)
	if err != nil {
		return Queue_Request{}, err
	}
	return request, err
}

func GetNextRequestInLine() {

}
