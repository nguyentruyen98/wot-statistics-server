package main

import (
	"fmt"
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

	_, err := postgres.NewDB(&config.DatabaseConfig{
		URL:             "postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable",
		MaxOpenConns:    25,
		MaxIdleConns:    25,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
	})

	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	appConfig := config.GetAppConfig()

	route.SetupRouter(router, appConfig)

	router.Run()

}
