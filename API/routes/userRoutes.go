package routes

import (
	"Nezha/API/controllers"
	"Nezha/API/helpers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(helpers.Authentication())
	incomingRoutes.GET("/users/apps", controllers.GetApps())
	incomingRoutes.GET("/app/connect/:id", controllers.Connect())
}
