package queue

import (
	"fmt"
	"encoding/json"
	"time"
	d "ECE49595_PROJECT/dock"
	"errors"
	"github.com/go-redis/redis"
)

var queue Queue //singleton implementation
var Queue_container_name string
func initQueue(api_conn_options, ssh_serv_conn_options *redis.Options) bool{
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
	return queue.ONLINE
}
func MakeQueue(api_conn_options, ssh_serv_conn_options *redis.Options) error{
	var make_err error
	if Queue_container_name != "" {
		d.StopOneContainer(Queue_container_name)
		d.RemoveOneContainer(Queue_container_name)
	}
	Queue_container_name, make_err =  d.CreateNewContainer(QUEUE_CONTAINER_IMAGE,QUEUE_CONTAINER_MACHINE_PORT, QUEUE_CONTAINER_PORT)
	if make_err != nil{
		fmt.Println("Could not make queue.",make_err)
		return make_err
	}
	if initQueue(api_conn_options, ssh_serv_conn_options){
		return nil
	}else{
		return errors.New("Could not init queue.")
	}
}

func makeQueueConnection(options *redis.Options, name string) (*redis.Client, error) {
	//begin client
	queueConn := redis.NewClient(options)
	//check if a container exists
	_, err := queueConn.Ping().Result()
	if err != nil {
		fmt.Println("Queue container could not be pinged.")
		return nil, err
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

func ShutDownQueue(force bool) int {
	if !QueueIsEmpty() && !force {
		return QUEUE_SHUTDOWN_FAIL_QUEUE_NOT_EMPTY
	}
	if EmptyQueue() != nil {
		return QUEUE_SHUTDOWN_FAIL_COULD_NOT_EMPTY_QUEUE
	}
	queue.API_CLI.Close()
	queue.SSH_SERV_CLI.Close()

	shutDownCheck := (d.StopOneContainer(Queue_container_name) == nil)
	if d.RemoveOneContainer(Queue_container_name) != nil{
		fmt.Println("Queue container stopped not removed.")
	}

	if !shutDownCheck {
		return QUEUE_SHUTDOWN_UNSUCCESSFUL
	}
	return QUEUE_SHUT_DOWN_SUCCESSFUL
}

func RestartQueue(force bool) int {
	if !QueueIsEmpty() && !force {
		return QUEUE_RESTART_FAIL_QUEUE_NOT_EMPTY
	}
	if EmptyQueue() != nil {
		return QUEUE_RESTART_FAIL_COULD_NOT_EMPTY_QUEUE
	}

	err:= d.RestartContainer(Queue_container_name); if err!=nil{
		return QUEUE_RESTART_UNSUCCESSFUL
	}
	if !initQueue(queue.API_CONN_OPTIONS, queue.SSH_SERV_CONN_OPTIONS) {
		return QUEUE_RESTART_UNSUCCESSFUL
	}
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
