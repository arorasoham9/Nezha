package API

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(Authentication())
	incomingRoutes.GET("/users/apps", GetApps())
	incomingRoutes.GET("/app/connect/:id", Connect())
	// incomingRoutes.GET("/ping", Connect())
}
