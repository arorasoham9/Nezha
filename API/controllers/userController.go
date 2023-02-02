package controllers

import (
	"fmt"
	"log"

	"github.com/arorasoham9/ECE49595_PROJECT/API/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		check = false
	}

	return check, msg
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		//var user models.User
		//var foundUser models.User
		fmt.Println("%+v", c)
	}
}
