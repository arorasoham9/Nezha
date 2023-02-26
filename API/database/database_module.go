package database

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseModule struct {
	MongoDB string
	client  *mongo.Client
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
