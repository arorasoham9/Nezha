package API

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppList struct {
	ID    primitive.ObjectID
	Email *string
	Apps  []*string
}

type App struct {
	ID    primitive.ObjectID `json:"Id,omitempty" bson:"_id,omitempty"`
	Users []*string
	Host  *string
	Port  int
	Name  *string
}
