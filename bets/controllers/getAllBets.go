package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SampleController Return the response of authenticated user.
// SampleController                godoc
// @Title 		loginController
// @Summary		request for authenticate users
// @Description  use for authenticate registered user with the help of email and password.
// @Tags         Bets
// @Produce      json
// @Param        username formData string  true  "registered username"
// @Param        password formData string  true  "password"
// @Success      200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/bets/getAllBets [get]
func GetAllBets(c *gin.Context) {
	fmt.Println("placebet Controller")
	c.JSON(http.StatusOK, gin.H{
		"message": "Controller Response",
	})

}
