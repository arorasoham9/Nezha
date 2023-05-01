package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/arorasoham9/ECE49595_PROJECT/API/database"
	"github.com/arorasoham9/ECE49595_PROJECT/API/helpers"
	"github.com/arorasoham9/ECE49595_PROJECT/API/models"
	"github.com/arorasoham9/ECE49595_PROJECT/API/queue"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/api/idtoken"
)

var db = database.DatabaseModule{}

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

// Rewrite to get email correctly from context
func GetApps() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("email")
		fmt.Println(email)
		appList, _ := db.GetApps(email)

		request := queue.Queue_Request{
			EMAIL:      email,
			CURRENT_IP: c.ClientIP(),
			CREATED_AT: time.Now().String(),
		}

		err := queue.SendToRedis(request, "mabaums")

		if err != nil {
			log.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.IndentedJSON(http.StatusOK, appList)
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
		count, err := db.GetEmailCount(*user.Email)
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

		clientID := "688933920583-vojm3og1kndonhvo6icej2r2q8a0la8b.apps.googleusercontent.com"

		payload, err := idtoken.Validate(context.Background(), *user.Token, clientID)
		if err != nil {
			log.Printf("Err %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error unauthorized."})
			return
		}
		//fmt.Print(payload.Claims["email"])
		claims := payload.Claims
		email := fmt.Sprintf("%v", claims["email"])
		log.Printf("Attempted login user %v", email)

		foundUser, err = db.FindUserByEmail("users", email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "login or passowrd is incorrect"})
			log.Printf("Invalid user: %v", *user.Email)
			return
		}

		token, _, err := helpers.GenerateAllTokens(*foundUser.Email) // TODO: Return refresh token.

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with JWT token creation"})
		}

		c.JSON(http.StatusOK, gin.H{"Token": token})
	}
}
