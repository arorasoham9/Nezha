package main

import (
	queue "Nezha/queue"
	"Nezha/API"
)

func main() {
	queue.BeginQueueOperation() 
	API.RunAPI()

}

