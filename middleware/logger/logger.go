package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// i want to write a logger middleware for gin framework
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()

		// Process request
		c.Next()

		// Stop timer
		elapsed := time.Since(t)

		// Log the request
		fmt.Printf("%s | %s | %s | %d | %s | %s | %s\n",
			t.Format("2006-01-02 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			c.Writer.Status(),
			c.Request.URL.Path,
			c.Errors.String(),
			elapsed,
		)
	}
}

// every request will be logged in this format
func FormatLogs(param gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"%s | %s | %s | %d | %s | %s | %s\n",
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		param.ClientIP,
		param.Method,
		param.StatusCode,
		param.Path,
		param.ErrorMessage,
		param.Latency.String(),
	)

}
