package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	//"user-athentication-golang/database"

	"github.com/arorasoham9/ECE49595_PROJECT/API/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both the detailed token and refresh token
func GenerateAllTokens(email string) (signedToken string, signedRefreshToken string, err error) {
	// TODO: Extract env var logic into helper Setting defaults if not present.
	tokenDuration, _ := strconv.ParseInt(os.Getenv("TOKEN_DURATION"), 10, 64)
	refreshTokenDuration, _ := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_DURATION"), 10, 64)
	claims := &SignedDetails{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(tokenDuration)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(refreshTokenDuration)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

// UpdateAllTokens renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	return
	/*
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var updateObj primitive.D

		updateObj = append(updateObj, bson.E{"token", signedToken})
		updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

		Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

		upsert := true
		filter := bson.M{"user_id": userId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		_, err := userCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)
		defer cancel()

		if err != nil {
			log.Panic(err)
			return
		}

		return */
}
