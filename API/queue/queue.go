package queue

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// Temporarily hosting queue request struct here for demo.
type Queue_Request struct {
	NAME       string `json:"name"`
	EMAIL      string `json:"email"`
	CURRENT_IP string `json:"current_ip"`
	LOCATION   string `json:"location"`
	CREATED_AT string `json:"created_at"`
	LASTSEEN   string `json:"last_seen"`
}

var client *redis.Client

func ConnectToRedis() error {
	address := os.Getenv("REDIS_HOST")
	if address == "" {
		return errors.New("No environment variable found: REDIS_HOST")
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		log.Warningf("No environment variable found: REDIS_PORT defaulting to: 6379")
		port = "6379"
	}

	// No redis password being set. TODO: add password env var.
	client = redis.NewClient(&redis.Options{
		Addr: address + ":" + port,
		DB:   0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	log.Infof("Connection Succes to Redis %v", pong)
	return nil
}

func SendToRedis(request Queue_Request, key string) error {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = client.Set(key, reqBytes, 0).Err() // No expiration at the moment.
	if err != nil {
		return err
	}

	return nil
}
