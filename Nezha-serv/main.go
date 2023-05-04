package main

import (
	"Nezha-serv/API"
	queue "Nezha-serv/queue"
	"Nezha-serv/ssh"
	"encoding/json"
	"fmt"
)

func main() {
	request := queue.Queue_Request{
		USERNAME:   API.DEFAULT_USERNAME,
		EMAIL:      API.DEFAULT_EMAIL,
		LOCATION:   API.DEFAULT_LOCATION,
		CREATED_AT: API.DEFAULT_CREATED_AT,
		DIAL_PORT:  API.DEFAULT_DIAL_PORT,
		BIND_PORT:  API.DEFAULT_BIND_PORT,
		HOST_ADDR:  API.DEFAULT_HOST_ADDR,
		HOST_PORT:  API.DEFAULT_HOST_PORT,
		PASSWORD:   API.DEFAULT_PASSWORD,
	}
	// queue.BeginQueueOperation()
	json_byte_request, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	ssh.SendOutConnection(json_byte_request, "test1")

}
