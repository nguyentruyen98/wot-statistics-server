package route

import (
	"wot-statistics-server/config"
	"wot-statistics-server/internal/infrastructure/postgres"
	"wot-statistics-server/wire"

	"github.com/gin-gonic/gin"
)

func NewTankRouter(group *gin.RouterGroup, db *postgres.DB, appConfig *config.AppConfig) {
	tc := wire.InitializeTankController(db, appConfig)

	// Tank routes
	tanks := group.Group("/tanks")
	{
		// GET endpoints
		tanks.GET("", tc.GetTanks)                  // Get all tanks with filters
		tanks.GET("/:id", tc.GetTankByID)           // Get single tank by ID
		tanks.GET("/wargaming", tc.GetTanksFromAPI) // Get from Wargaming API

		// POST endpoints
		tanks.POST("", tc.CreateTank)                 // Create new tank
		tanks.POST("/import", tc.ImportTanksFromJSON) // Import from JSON file

		// PUT endpoints
		tanks.PUT("/:id", tc.UpdateTank) // Update tank

		// DELETE endpoints
		tanks.DELETE("/:id", tc.DeleteTank) // Delete tank
	}
}
