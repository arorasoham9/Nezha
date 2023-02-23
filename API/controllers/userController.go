package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arorasoham9/ECE49595_PROJECT/API/database"
	"github.com/arorasoham9/ECE49595_PROJECT/API/helpers"
	"github.com/arorasoham9/ECE49595_PROJECT/API/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

func GetApps() gin.HandlerFunc {
	return func(c *gin.Context) {
		Apps := []string{"App 1", "App 2", "App 3"}
		c.IndentedJSON(http.StatusOK, Apps)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		validationErr := validate.Struct(user)
		defer cancel()
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Print(count)

		//defer cancel()
		/*
			user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.ID = primitive.NewObjectID()
			user.User_id = user.ID.Hex()
			token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
			user.Token = &token
			user.Refresh_token = &refreshToken */
		fmt.Println(*user.Email)
		//var foundUser models.User

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Printf("Attempted login user %v", *user.Email)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "login or passowrd is incorrect"})
			log.Printf("Invalid user: %v", *user.Email)
			return
		}

		token, _, err := helpers.GenerateAllTokens(*foundUser.Email) // TODO: Return refresh token.

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with JWT token creation"})
		}

		c.JSON(http.StatusOK, token)
	}
}
