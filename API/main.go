package main

import (
	"time"

	"Nezha/API/helpers"
	"Nezha/API/queue"
	"Nezha/API/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := queue.ConnectToRedis()
	if err != nil {
		log.Fatalf("%v", err)
	}

	port := helpers.GetPort()
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
