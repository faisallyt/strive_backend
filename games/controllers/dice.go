package controllers

import (
	"math/rand"
	"net/http"
	dice "strive_go/games/services/gamePlay/dice"
	"strive_go/games/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func randomNumberGeneration() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2)
}

// GetBookByISBN locates the book whose ISBN value matches the isbn
// GetBookByISBN                godoc
// @Summary      Get the result of bet
// @Description  Returns the amount of money that the user won or lost
// @Tags         Dice
// @Produce      json
// @Param        betAmount query string true "BetAmount"
// @Param        rollOver query string true "RollOver"
// @Param        username query string true "Username"
// @Success      200  {object}  string  "ok"
// @Router       /api/v1/games/dice [get]
func DiceController(c *gin.Context) {

	betAmount := c.Query("betAmount")
	rollOver := c.Query("rollOver")
	// username := c.Query("username")//TODO: use this username to get user details from database

	result, err := dice.DiceOutput(rollOver, betAmount)
	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Successfully Played", result)

}
