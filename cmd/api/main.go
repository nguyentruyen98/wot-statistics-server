package main

import (
	"fmt"
	"log"
	"time"
	"wot-statistics-server/api/route"
	"wot-statistics-server/config"
	"wot-statistics-server/internal/infrastructure/postgres"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Cấu hình CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Connect to database
	db, err := postgres.NewDB(&config.DatabaseConfig{
		URL:             "postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable",
		MaxOpenConns:    25,
		MaxIdleConns:    25,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	fmt.Println("✅ Database connected successfully")

	// Get app config
	appConfig := config.GetAppConfig()

	// Setup routes
	route.SetupRouter(router, db, appConfig)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "WoT Statistics Server is running",
		})
	})

	// Start server
	fmt.Println("🚀 Server starting on :8080")
	router.Run(":8080")
}
