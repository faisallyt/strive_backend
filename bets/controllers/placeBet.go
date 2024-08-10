package controllers

import (
	"fmt"
	"net/http"
	"strive_go/payment/utils"

	"github.com/gin-gonic/gin"
)

// PlaceBetController places a bet for game
// PlaceBetController                godoc
// @Title 		PlaceBetController
// @Summary		request for place the bet for game
// @Description  use for place the bet using username and amount
// @Tags         Bet
// @Produce      json
// @Param        username query string  true  "registered username"
// @Param        amount formData float64  true  "amount to place"
// @Success      200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/bet/placebet [post]
func PlaceBetController(c *gin.Context) {
	fmt.Println("placebet Controller")
	c.JSON(http.StatusOK, gin.H{
		"message": "Controller Response",
	})
	utils.SendSuccessResponse(c, http.StatusOK, "bet placed successfully", "ok")
}
