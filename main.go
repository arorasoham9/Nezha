package main

import (
	queue "Nezha/queue"
)

func main() {
	// address := os.Getenv("REDIS_HOST")
	// if address == "" {
	// 	log.Fatal("No environment variable found: REDIS_HOST")
	// }

	// port := os.Getenv("REDIS_PORT")
	// if port == "" {
	// 	log.Warningf("No environment variable found: REDIS_PORT defaulting to: 6379")
	// 	port = "6379"
	// }

	// log.Infof("Starting Queue connecting to Redis: %v:%v", address, port)

	queue.BeginQueueOperation() 
	//this is a dummy change, please ignore
}
