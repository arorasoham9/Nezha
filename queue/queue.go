package queue

import (
	// "fmt"
	"os"
	"time"
	"github.com/go-redis/redis"
)

func handleErr(err  error, cliname string){
	if err != nil{
		os.Setenv("QUEUE_AVAILABLE", UNSET )
		os.Setenv(cliname+"_ONLINE", UNSET )
		
	} else{
		os.Setenv("QUEUE_AVAILABLE", SET )
		os.Setenv(cliname+"_ONLINE", SET )
	}
}

func MakeQueue(api_conn_options, ssh_serv_conn_options *redis.Options) Queue {
	api_connection, err1 := MakeQueueConnection(api_conn_options, API_Q_CLI)
	ssh_serv_connection, err2:= MakeQueueConnection(ssh_serv_conn_options, SSH_Q_CLI )
	queue := Queue{
		API_CLI: api_connection,
		SSH_SERV_CLI: ssh_serv_connection,
		CREATED: time.Now(),
		ONLINE: err1 !=nil || err2 !=nil,
		API_CONN_OPTIONS: api_conn_options,
		SSH_SERV_CONN_OPTIONS: ssh_serv_conn_options,
	}
	return queue
}
func MakeQueueConnection(options *redis.Options, name string ) (*redis.Client, error)   {
	//begin client
	queueConn := redis.NewClient(options)
	//check if a container exists
	_, err := queueConn.Ping().Result()
	handleErr(err, name)
	return queueConn, err
}


func CheckAlive(queue Queue)Queue{
	_, err1 := queue.API_CLI.Ping().Result()
	_, err2 := queue.SSH_SERV_CLI.Ping().Result()
	if err1 !=nil || err2 !=nil {
		queue.API_CLI.Close()
		queue.SSH_SERV_CLI.Close()
		queue = MakeQueue(queue.API_CONN_OPTIONS,queue.SSH_SERV_CONN_OPTIONS)
	}
	return queue
}

func handleRequest(queue Queue, EXIT_CODE int){
	
	//

}

func parseJSON(){

}


func shutDownQueue(){

}
func raiseInterrupt(){
	//little unsure how to make this 

}


func clearInterrupt(){

}