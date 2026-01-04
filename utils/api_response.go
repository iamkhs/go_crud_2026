package utils

import (
	"github.com/gin-gonic/gin"
)

// ApiResponse represents the standard structure for all API responses
type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendSuccessResponse sends a JSON response with a success status (200, 201, etc.)
func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := ApiResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}

// SendErrorResponse sends a JSON response with an error status (400, 404, 500, etc.)
func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	response := ApiResponse{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	}
	c.JSON(statusCode, response)
}
