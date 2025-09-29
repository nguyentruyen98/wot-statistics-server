package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the request method and path
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

		// Process the request
		c.Next()

		// Log the response status
		log.Printf("Response: %d", c.Writer.Status())
	}
}
