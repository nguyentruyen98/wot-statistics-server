package route

import (
	"wot-statistics-server/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, appConfig *config.AppConfig) {

	publicRouter := router.Group("/api")

	NewTankRouter(publicRouter, appConfig)
}
