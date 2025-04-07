package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/TechnoDiktator/ginTutorial/utils/models/sharedpackets"
)

func RouteRequestHandleTemp(c *gin.Context, localServiceInit *sharedpackets.LocalServiceInit) {
	// Handle the request here
	// You can use the context passed to this function to manage request-scoped values
	// and cancellation signals.

	// Example: ctx.Value("key") to get a value from the context
	// Example: ctx.Done() to check if the context is done
	// Example: ctx.Err() to check if there was an error with the context
	// Example: ctx.WithTimeout() to create a new context with a timeout

	var requestPacket map[string]interface{}

	

	if err := c.ShouldBindJSON(&requestPacket); err != nil {
		logrus.Errorf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
}
