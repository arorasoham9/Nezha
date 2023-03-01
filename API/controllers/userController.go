package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arorasoham9/ECE49595_PROJECT/API/database"
	"github.com/arorasoham9/ECE49595_PROJECT/API/helpers"
	"github.com/arorasoham9/ECE49595_PROJECT/API/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var db = database.DatabaseModule{}

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
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
		//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		validationErr := validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, err := db.GetEmailCount("users", *user.Email)
		//count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Print(count)

		fmt.Println(*user.Email)

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

		foundUser, err := db.FindUserByEmail("users", *user.Email)
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
