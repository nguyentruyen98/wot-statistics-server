package route

import (
	"wot-statistics-server/config"
	"wot-statistics-server/internal/infrastructure/postgres"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *postgres.DB, appConfig *config.AppConfig) {
	publicRouter := router.Group("/api")

	NewTankRouter(publicRouter, db, appConfig)
}
