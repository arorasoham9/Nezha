package API

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/idtoken"
)

// var db = DatabaseModule{}

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

// Rewrite to get email correctly from context
func GetApps() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("email")
		// name := c.GetString("name")
		fmt.Println(email)
		appList, _ := db.GetApps(email)

		c.IndentedJSON(http.StatusOK, appList)
	}
}

func Connect() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id") // Do something with app Id
		if !ok {
			log.Errorf("Connection requires id")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		email := c.GetString("email")
		name := c.GetString("name")

		request := queue.Queue_Request{
			USERNAME:   DEFAULT_USERNAME,
			EMAIL:      DEFAULT_EMAIL,
			LOCATION:   DEFAULT_LOCATION,
			CREATED_AT: DEFAULT_CREATED_AT,
			DIAL_PORT:  DEFAULT_DIAL_PORT,
			BIND_PORT:  DEFAULT_BIND_PORT,
			HOST_ADDR:  DEFAULT_HOST_ADDR,
			HOST_PORT:  DEFAULT_HOST_PORT,
			PASSWORD:   DEFAULT_PASSWORD,
		}
		queue.QueueLck.Lock()
		err := queue.AddRequestToQueue(email+"::"+name+"::"+id, request)
		queue.QueueLck.Unlock()
		if err != nil {
			log.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user User
		var foundUser *User

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
		name := claims["name"].(string)

		log.Printf("Attempted login user %v", email)

		foundUser, err = db.FindUserByEmail(email)
		if err != nil {
			log.Infof("Creating new user %v", email)
			foundUser, err = db.AddUser(email, name, false)
			if err != nil {
				log.Errorf("Could not create new user")
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		token, _, err := GenerateAllTokens(email, name, foundUser.IsAdmin) // TODO: Return refresh token.

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with JWT token creation"})
		}

		c.JSON(http.StatusOK, gin.H{"Token": token})
	}
}
