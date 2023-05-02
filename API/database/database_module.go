// Package database
package database

import (
	"context"
	"fmt"
	"time"

	"Nezha/API/helpers"
	"Nezha/API/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// getApps returns a Go Array of Strings and an error
// The function checks the database for a set of apps that a particular user has access too.
func (d DatabaseModule) GetApps(email string) ([]models.App, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	appCollection := d.openCollection("apps")
	defer cancel()
	cursor, err := appCollection.Find(ctx, bson.M{"users": email})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var foundApp models.App
	appList := []models.App{}
	for cursor.Next(ctx) {
		err = cursor.Decode(&foundApp)
		if err != nil {
			log.Errorf("Error decoding found app %v", err)
		} else {
			appList = append(appList, foundApp)
		}
	}

	return appList, err
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
func (d DatabaseModule) FindUserByEmail(email string) (*models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.openCollection("users")
	defer cancel()
	res := userCollection.FindOne(ctx, bson.M{"email": email})
	var foundUser models.User
	err := res.Decode(&foundUser)
	return &foundUser, err
}

// AddUser
func (d DatabaseModule) AddUser(email string, name string, isAdmin bool) (*models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userCollection := d.openCollection("users")
	defer cancel()
	user := bson.M{
		"email":   email,
		"name":    name,
		"isAdmin": isAdmin,
	}
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return d.FindUserByEmail(email)
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
