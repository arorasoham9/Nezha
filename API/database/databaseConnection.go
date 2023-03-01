package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arorasoham9/ECE49595_PROJECT/API/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function name inputs/outputs/basic description

func DBinstance() *mongo.Client {
	MongoDb := helpers.GetMongoURL()

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB") // TODO: Change to log

	return client
}

var Client *mongo.Client = DBinstance()

/*
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("test").Collection(collectionName)

	return collection
} */
