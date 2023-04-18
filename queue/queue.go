package queue

import (
	d "ECE49595_PROJECT/dock"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var queue Queue //singleton implementation
var Queue_container_name string
var CurrentKey string
var MasterWorker QueueWorker
var slaveWorkers []QueueWorker
var condvar *sync.Cond

func initQueue(api_conn_options, ssh_serv_conn_options *redis.Options) bool{
	api_connection, err1 := makeQueueConnection(api_conn_options, API_Q_CLI)
	ssh_serv_connection, err2 := makeQueueConnection(ssh_serv_conn_options, SSH_Q_CLI)
	condvar = &sync.Cond{L:&sync.Mutex{}}
	queue = Queue{
		API_CLI:               api_connection,
		SSH_SERV_CLI:          ssh_serv_connection,
		CREATED:               time.Now(),
		ONLINE:                err1 == nil && err2 == nil,
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
// func GetRequestNextInLine()(Queue_Request, error){

// }
func InitWorker(worker *QueueWorker, _ID int, isMaster bool){
	worker.cond = condvar
	worker.ID = _ID
	worker.CREATED = time.Now()
	worker.SERVED = 0
	worker.MASTER = isMaster
}
//Right now the plan is to entertain one request at a time but open connections to the users via go routines. But this function will also be called by a goroutine which constantly checks
//the appearance of new requests and opens connections.
func BeginWork(worker *QueueWorker) {
	// failcount :=0
	for {
		fmt.Println("Do I come here?")
		worker.cond.L.Lock()
		if QueueIsEmpty(){
			worker.cond.Wait()
		}
		//now send out a connection.
		// request, err := GetRequestNextInLine()
		// if err != nil && failcount > MAX_GET_NEXT_REQUEST_FAIL{
		// 	failcount++
		// 	continue
		// }
		worker.cond.L.Unlock()
		//take care of request
		if worker.SERVED >5 {
			fmt.Println("slave with ID", worker.ID,"exiting")
			break
		}


		//increment num of sessions helped
		worker.SERVED++
		
		
	}

}
func StartAllWorkers(maxSlaves, masterID int){
	InitWorker(&MasterWorker, masterID, true)
	fmt.Println("Do I come herggge?", MasterWorker.ID, MasterWorker.SERVED, MasterWorker.MASTER)
	go BeginWork(&MasterWorker)
	// for i:=0; i< maxSlaves; i++{
	// 	var slave QueueWorker
	// 	InitWorker(&slave, masterID+i+1, false)
	// 	slaveWorkers = append(slaveWorkers, slave)
	// 	go BeginWork(&slave)
	// }
}

func BeginQueueOperation(api_conn_options, ssh_serv_conn_options *redis.Options, maxSlaves, masterID int ){
	d.InitDock()
	err :=MakeQueue(api_conn_options, ssh_serv_conn_options); if err != nil{
		fmt.Println(err)
		os.Exit(HARD_EXIT)
	}
	StartAllWorkers(maxSlaves, masterID)
}