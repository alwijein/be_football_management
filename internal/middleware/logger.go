package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"log"
)

// Logger middleware logs HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Get status code
		statusCode := c.Writer.Status()

		// Log format
		log.Printf("[%s] %s %s %d %v",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			statusCode,
			latency,
		)
	}
}
