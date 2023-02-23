package main

import (
	"os"

	"github.com/arorasoham9/ECE49595_PROJECT/API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.Run(":" + port)
}
