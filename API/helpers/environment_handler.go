package helpers

import (
	"log"
	"os"
	"strconv"
)

func checkExist(varName string) (string, bool) {
	envVar, found := os.LookupEnv(varName)
	return envVar, found
}

func GetSecretKey() string {
	SECRET_KEY, found := checkExist("SECRET_KEY")
	if !found {
		log.Panic("No SECRET KEY found, can not encode JWT")
	}
	return SECRET_KEY
}

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

func GetMongoURL() string {
	MongoDb, found := checkExist("MONGODB_URL")
	if !found {
		log.Panic("No DB URL found")
		//MongoDb = "mongodb://localhost:27017"
	}
	return MongoDb
}
