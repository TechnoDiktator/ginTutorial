package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/TechnoDiktator/ginTutorial/utils/constants"
	"github.com/TechnoDiktator/ginTutorial/utils/models/dbmodels"
	"github.com/TechnoDiktator/ginTutorial/utils/models/requestpackets"
	"github.com/TechnoDiktator/ginTutorial/utils/models/responsepackets"
	"github.com/TechnoDiktator/ginTutorial/utils/models/sharedpackets"
	"github.com/TechnoDiktator/ginTutorial/utils/validators"
)

func RouteRequestHandleTemp(c *gin.Context, localServiceInit *sharedpackets.LocalServiceInit) {
	// Handle the request here
	// You can use the context passed to this function to manage request-scoped values
	// and cancellation signals.

	// Example: ctx.Value("key") to get a value from the context
	// Example: ctx.Done() to check if the context is done
	// Example: ctx.Err() to check if there was an error with the context
	// Example: ctx.WithTimeout() to create a new context with a timeout

	var requestPacket requestpackets.RequestPacketTypeOne

	if err := c.ShouldBindJSON(&requestPacket); err != nil {
		logrus.Errorf("Error binding JSON: %v", err)
		//c.JSON(constants.StateInvalidRequestPacket.HttpsStatusCode, gin.H{"error": "Invalid JSON"})

		c.JSON(constants.StateInvalidRequestPacket.HttpsStatusCode, responsepackets.CommanErrorResponse{
			Success: constants.StateInvalidRequestPacket.Success,
			Message: constants.StateInvalidRequestPacket.Message,
			Code:    constants.StateInvalidRequestPacket.ErrorCode,
			Errors:  err.Error(),
			Data:    nil,
		})

		valid, typeErrors, emptyErrors, fetchErrors, dataMatchingError := validators.ValidateRequestForTempObject(requestPacket)
		if !valid {
			var errors []string
			if len(typeErrors) > 0 {
				errors = append(errors, typeErrors...)
			}
			if len(emptyErrors) > 0 {
				errors = append(errors, emptyErrors...)
			}
			if len(fetchErrors) > 0 {
				errors = append(errors, fetchErrors...)
			}
			if len(dataMatchingError) > 0 {
				errors = append(errors, dataMatchingError...)
			}
			c.JSON(constants.StateValidationFailed.HttpsStatusCode, responsepackets.CommanErrorResponse{
				Success: constants.StateValidationFailed.Success,
				Message: constants.StateValidationFailed.Message,
				Errors:  errors,
				Code:    constants.StateValidationFailed.ErrorCode,
				Data:    nil,
			})
			return
		}

		var data *dbmodels.Temp
		jsondata, err := json.Marshal(data)

		RawMessage := json.RawMessage(jsondata)

		if err != nil {
			logrus.Errorf("Error marshalling data: %v", err)
			c.JSON(http.StatusInternalServerError, responsepackets.CommanErrorResponse{
				Success: false,
				Message: "Internal server error",
				Code:    500,
				Data:    nil,
				Errors:  err.Error(),
			})
			return
		}
		// Return a Success response
		c.JSON(http.StatusAccepted, responsepackets.CommonResponsePacket{
			Success: true,
			Code:    "200",
			Message: "Request processed successfully",
			Data:    RawMessage,
		})
		return

	}

	//maybe some validator should also be put here
	// Process the requestPacket as needed

}
