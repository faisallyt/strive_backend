package utils

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := ApiResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	c.JSON(statusCode, response)
}
