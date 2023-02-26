package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseModule struct {
	client *mongo.Client
}

func (d DatabaseModule) getClient() *mongo.Client {
	if d.client == nil {
		d.client = DBinstance()
	}
	return d.client
}

func (d DatabaseModule) OpenCollection(collectionName string) *mongo.Collection {
	cl := d.getClient()
	var collection *mongo.Collection = cl.Database("test").Collection(collectionName)
	return collection
}

func (d DatabaseModule) GetEmailCount(collectionName string, email string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.OpenCollection(collectionName)
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": email})
	defer cancel()
	if err != nil {
		return -1, err
	}
	return count, err
}

func (d DatabaseModule) FindEmail(collectionName string, email string) *mongo.SingleResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.OpenCollection(collectionName)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.M{"email": email})
	return res
}

func (d DatabaseModule) AddUser(email string) {
	return
}

func (d DatabaseModule) CreateCollection(collectionName string) {
	return
}
