package API

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"

)

func RunAPI() {


	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //bad idea
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	AuthRoutes(router)
	UserRoutes(router)
	
	router.Run(":" + API_PORT)
}
