package logger

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


//every request will be logged in this format
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