package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppList struct {
	ID    primitive.ObjectID
	Email *string
	Apps  []*string
}
