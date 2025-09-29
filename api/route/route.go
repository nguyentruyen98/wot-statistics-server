package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	publicRouter := router.Group("/api")

	NewTankRouter(publicRouter)
	// privateRouter := router.Group("/api").Use(middleware.LoggerMiddleware())

	// privateRouter.GET("/private", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "This is a private endpoint",
	// 	})
	// })

	// publicRouter.GET("/public", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "This is a public endpoint",
	// 	})
	// })
}
