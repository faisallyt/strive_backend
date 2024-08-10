package utils

import "github.com/gin-gonic/gin"

type ApiError struct {
	StatusCode int    `json:"Status_code"`
	Message    string `json: "message"`
}

func SendApiError(c *gin.Context, statusCode int, message string) {
	apiError := ApiError{
		StatusCode: statusCode,
		Message:    message,
	}
	c.JSON(statusCode, apiError)
	c.Abort()
}
