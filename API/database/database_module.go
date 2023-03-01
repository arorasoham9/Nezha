package database

import (
	"context"
	"log"
	"time"

	"github.com/arorasoham9/ECE49595_PROJECT/API/models"
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

func (d DatabaseModule) openCollection(collectionName string) *mongo.Collection {
	db := d.openDatabse("test")
	var collection *mongo.Collection = db.Collection(collectionName)
	return collection
}

func (d DatabaseModule) GetEmailCount(collectionName string, email string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.openCollection(collectionName)
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": email})
	defer cancel()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return count, err
}

func (d DatabaseModule) FindUserByEmail(collectionName string, email string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.openCollection(collectionName)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.M{"email": email})
	var foundUser models.User
	err := res.Decode(&foundUser)
	return foundUser, err
}

func (d DatabaseModule) AddUser(email string) {
	return
}

func (d DatabaseModule) openDatabse(dbname string) *mongo.Database {
	cl := d.getClient()
	db := cl.Database(dbname)
	return db
}

func (d DatabaseModule) CreateCollection(collectionName string) {
	return
}
