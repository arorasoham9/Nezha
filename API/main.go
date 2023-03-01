package main

import (
	"github.com/arorasoham9/ECE49595_PROJECT/API/helpers"
	"github.com/arorasoham9/ECE49595_PROJECT/API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := helpers.GetPort()
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
