package API

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID
	Email    *string
	Password *string
	Token    *string
	IsAdmin  bool
}
