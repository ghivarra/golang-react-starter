package customError

import (
	"backend/config/environment"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Build error message based on "ENV" value in .env file
func buildMessage(message string, err error) string {
	if environment.ENV == "development" {
		message += fmt.Sprintf(" Reason: %v", err)
	}

	return message
}

// Build JSON context message based on "ENVIRONMENT"
func buildJSON(message string, err error) map[string]any {
	json := gin.H{
		"status":  "error",
		"message": message,
	}

	// add error as strings
	if environment.ENV == "development" {
		json["data"] = fmt.Sprintf("%v", err)
	}

	// return
	return json
}

// Send Error Log to the console.
func SendErrorLog(message string, err error) {
	fmt.Println(buildMessage(message, err))
}

// Send Error Response to the interface and Error Log to the console.
func SendErrorResponse(c *gin.Context, status int, message string, err error) {

	// send error log, use goroutine
	go SendErrorLog(message, err)

	// build message
	message = buildMessage(message, err)
	jsonData := buildJSON(message, err)

	// send to interface
	c.JSON(status, jsonData)
}
