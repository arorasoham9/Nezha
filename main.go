package main

import (
	"ECE49595_PROJECT/queue"
	"fmt"
	"os"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func main() {
	address := os.Getenv("REDIS_HOST")
	if address == "" {
		log.Fatal("No environment variable found: REDIS_HOST")
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		log.Warningf("No environment variable found: REDIS_PORT defaulting to: 6379")
		port = "6379"
	}

	log.Infof("Starting Queue connecting to Redis: %v:%v", address, port)
	options := redis.Options{
		Addr:     fmt.Sprintf("%v:%v", address, port),
		Password: "",
		DB:       0,
	}
	queue.BeginQueueOperation(&options, &options, 5, 5000) 
	//this is a dummy change, please ignore
}
