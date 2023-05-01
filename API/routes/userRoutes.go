package routes

import (
	"github.com/arorasoham9/ECE49595_PROJECT/API/controllers"
	"github.com/arorasoham9/ECE49595_PROJECT/API/helpers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(helpers.Authentication())
	incomingRoutes.GET("/users/apps", controllers.GetApps())
	incomingRoutes.GET("/app/connect/:id", controllers.Connect())
}
