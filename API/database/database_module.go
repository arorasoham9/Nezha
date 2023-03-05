// Package database
package database

import (
	"context"
	"log"
	"time"

	"github.com/arorasoham9/ECE49595_PROJECT/API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DatabaseModule represents a struct that holds a mongoClient allowing for Collection and Database access with extra error handling.
type DatabaseModule struct {
	client *mongo.Client
}

// getClient returns a Mongo Client from the Database instance
func (d DatabaseModule) getClient() *mongo.Client {
	if d.client == nil {
		d.client = DBinstance()
	}
	return d.client
}

// openCollection returns a Mongo Collection from a specific database with name collectionName
func (d DatabaseModule) openCollection(collectionName string) *mongo.Collection {
	db := d.openDatabse("test")
	var collection *mongo.Collection = db.Collection(collectionName)
	return collection
}

func (d DatabaseModule) GetApps(email string) ([]*string, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	appCollection := d.openCollection("apps")
	defer cancel()
	res := appCollection.FindOne(ctx, bson.M{"email": email})
	var foundApps models.AppList
	err := res.Decode(&foundApps)
	return foundApps.Apps, err
}

// GetEmailCount returns int64,error, the int represents the count of a certain email, and error if there was an error counting documents
// This function may become deprecated soon, was mainly used for testing and setting up API code.
func (d DatabaseModule) GetEmailCount(email string) (int64, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.openCollection("users")
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": email})
	defer cancel()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return count, err
}

// FindUserByEmail returns a Models User and error.
// The models User represents the user if they were fond
// Erorr represents any errors encountered
func (d DatabaseModule) FindUserByEmail(collectionName string, email string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.openCollection(collectionName)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.M{"email": email})
	var foundUser models.User
	err := res.Decode(&foundUser)
	return foundUser, err
}

// AddUser
func (d DatabaseModule) AddUser(email string) {
	return
}

// openDatabase returns a Mongo Database
// The Database of name dbname is opened using the client specified in the struct.
func (d DatabaseModule) openDatabse(dbname string) *mongo.Database {
	cl := d.getClient()
	db := cl.Database(dbname)
	return db
}

// CreateCollection
func (d DatabaseModule) CreateCollection(collectionName string) {
	return
}
