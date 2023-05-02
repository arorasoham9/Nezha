package queue

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)
const(
//add more stuff here as you need
	HARD_EXIT = 555
	SOFT_EXIT = 556
	QUEUE_UNAVAILABLE = 557;
	QUEUE_CONTAINERISED_UNAVAILABLE = 558
	QUEUE_CONTAINER_UNAVAILABLE = 559
	QUEUE_CREATED = 560 
	REDIS_AVAILABLE = 561
	REDIS_UNAVAILABLE = 562 
	SET = "true"
	UNSET = "false"
	API_Q_CLI = "API_Q_CLI"
	SSH_Q_CLI = "SSH_Q_CLI"
	EXIT_SUCCESS = 0
	EXIT_FAILURE = 1
	QUEUE_RESTART_FAIL_QUEUE_NOT_EMPTY = 5453
	QUEUE_RESTART_FAIL_COULD_NOT_EMPTY_QUEUE = 64564
	QUEUE_SHUTDOWN_FAIL = 64564
	QUEUE_RESTART_SUCCESSFUL = 545562
	QUEUE_RESTART_UNSUCCESSFUL = 7434377545562
	QUEUE_SHUTDOWN_UNSUCCESSFUL = 583283838
	QUEUE_SHUT_DOWN_SUCCESSFUL = 82382
	TTL = 1000000000
	QUEUE_CONTAINER_IMAGE = "redis/redis-stack"
	QUEUE_CONTAINER_PORT = "6379"
	QUEUE_CONTAINER_MACHINE_PORT = "6379"
	QUEUE_SHUTDOWN_FAIL_QUEUE_NOT_EMPTY = 545357489375948353
	QUEUE_SHUTDOWN_FAIL_COULD_NOT_EMPTY_QUEUE = 645645693543457888
	MAX_GET_NEXT_REQUEST_FAIL = 10
	QUEUE_FRONTEND_WORKER = 1
	QUEUE_BACKEND_WORKER = 2
	QUEUE_MASTER_WORKER = 3
	QUEUE_MAX_SLAVES = 10
	QUEUE_MASTER_ID = 5435346363576999999
	QUEUE_LOCALHOST = "127.0.0.1"
	QUEUE_DB_ID = 0
	QUEUE_DB_PASSWORD = ""
	QUEUE_START_CONTAINER = false
	QUEUE_KILL_SIGNAL_ENV_VAR = "QUEUE_KILL_SIGNAL_ENV_VAR"
	QUEUE_KILL_SET = "KILL_IT"
	QUEUE_KILL_UNSET = "DONT_KILL_IT"
	QUEUE_BAD_REQUEST = "BAD_REQUEST"
)


//add more stuff here as you need
type Queue_Request struct {
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
//add more stuff here as you need
type Location struct{

	//TBD

}
type QueueWorker struct{
	CondVarEmpty *sync.Cond
	CondVarAvailable *sync.Cond
	Lck *sync.Mutex
	ID int
	CREATED time.Time
	SERVED int
	MASTER bool
	SIDE int
	ONLINE bool
	API_CLI	*redis.Client 
	SSH_SERV_CLI *redis.Client  
}
type Queue struct{
	API_CLI	*redis.Client 
	SSH_SERV_CLI *redis.Client  
	CREATED time.Time 
	ONLINE bool 
	SSH_SERV_CONN_OPTIONS *redis.Options 
	API_CONN_OPTIONS *redis.Options 
}

