package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request
		method := c.Request.Method
		path := c.Request.URL.Path
		startTime := time.Now()

		fmt.Printf("Incoming request: %s %s\n", method, path)

		// Process the request
		c.Next()

		// Log the outgoing response
		statusCode := c.Writer.Status()
		duration := time.Since(startTime)

		fmt.Printf("Outgoing response: %s %s - Status: %d - Duration: %v\n", method, path, statusCode, duration)
	}
}
