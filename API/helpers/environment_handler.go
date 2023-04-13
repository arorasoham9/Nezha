// Package helpers implements different helper functions used in other areas of the project
// helpers/environemnt_handler.go impelements handlers used to interact with env vars.
package helpers

import (
	"log"
	"os"
	"strconv"
)

var TESTING_KEY string = ""

// checkExist returns a Go string and bool representing the environment variable and whether it was found
func checkExist(varName string) (string, bool) {
	envVar, found := os.LookupEnv(varName)
	return envVar, found
}

// GetPort returns a Go string that represents the PORT that will be set up for the API
// It's default value is 8000
func GetPort() string {
	port, found := checkExist("PORT")
	if !found {
		port = "8000"
	}
	return port
}

// GetSecretKey returns a Go String that represents the SECRET_KEY used in JWT enconding
func GetSecretKey() string {
	if TESTING_KEY != "" {
		return TESTING_KEY
	}
	key, found := checkExist("SECRET_KEY")
	if !found {
		log.Panic("No SECRET KEY found, can not encode JWT")
	}
	return key
}

// GetTokenDuration returns a Go int64 that represents how long the original token should last
func GetTokenDuration() int64 {
	duration := int64(24)
	durationS, found := checkExist("TOKEN_DURATION")
	if found {
		durationI, err := strconv.ParseInt(durationS, 10, 64)
		if err == nil {
			duration = int64(durationI)
		}
	}
	return duration
}

// GetRefreshTokenDuration returns a go int64 that represents how long the refresh token should last
func GetRefreshTokenDuration() int64 {
	duration := int64(196)
	durationS, found := checkExist("REFRESH_TOKEN_DURATION")
	if found {
		durationI, err := strconv.ParseInt(durationS, 10, 64)
		if err == nil {
			duration = int64(durationI)
		}
	}
	return duration
}

// GetMongoURL returns a Go string representing the URL for connecting to the MongoDB
func GetMongoURL() string {
	MongoDb, found := checkExist("MONGODB_URL")
	if !found {
		log.Panic("No DB URL found")
		//MongoDb = "mongodb://localhost:27017"
	}
	return MongoDb
}
