package mines

import (
	"net/http"
	"strive_go/auth/utils"
	"strive_go/db/functions/mines"
	gameplay "strive_go/games/services/gamePlay"

	"github.com/gin-gonic/gin"
)

type NextBetData struct {
	Box_No int64 `json:"box_no" binding:"required" description:"box number"`
}

func MinesNextController(c *gin.Context) {
	username, exists := c.Get("username")

	if !exists {
		utils.SendApiError(c, http.StatusUnauthorized, "User is not authorized")
		return
	}
	isActive, gameID, err := mines.CheckActiveGame(username.(string))

	if !isActive && err != nil {
		utils.SendApiError(c, http.StatusNotFound, "No active game found")
		return
	}

	if isActive && gameID == nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Error getting game ID. More than one game")
		return
	}

	activeGame, err := mines.FindCurrentMinesData(gameID)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Error getting game data")
		return
	}

	var NextBet NextBetData

	if err := c.ShouldBindJSON(&NextBet); err != nil {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Check if the box number is valid

	if NextBet.Box_No < 0 || NextBet.Box_No > 24 {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid box number")
		return
	}

	//check if the box number is already in the list

	if mines.IsBoxChosenAlready(activeGame.UserChosenBoxes, NextBet.Box_No) {
		utils.SendApiError(c, http.StatusBadRequest, "Box number already chosen")
		return
	}

	//check if the number of  remaining boxes is equal to the number of mines

	var remainingChances = 25 - activeGame.ChosenCount

	if remainingChances <= activeGame.MinesCount {
		err := mines.DebitUserBalance(username.(string), activeGame)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error making user lose")
			return
		}

		err = mines.AddBoxesWithMines(activeGame, NextBet.Box_No, activeGame.MinesCount)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error adding mines to boxes")
			return
		}
		err = mines.EndMinesGame(activeGame)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error updating game status")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, "User Lost ", activeGame.BoxesWithMines)
		return
	}

	//calculate multiplier
	multiplier := 99 / (4 * ((25 - activeGame.ChosenCount) - (activeGame.MinesCount)))

	result := gameplay.ComputeResult(float64(multiplier), activeGame.WinningTillNow, username.(string))

	if !result {
		err := mines.DebitUserBalance(username.(string), activeGame)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error making user lose")
			return
		}
		err = mines.AddBoxesWithMines(activeGame, NextBet.Box_No, activeGame.MinesCount)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error adding mines to boxes")
			return
		}
		err = mines.EndMinesGame(activeGame)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error updating game status")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, "ooops User Lost ", activeGame.BoxesWithMines)
		return
	} else {
		//is user wins
		err := mines.CreditUserBalance(float64(multiplier), activeGame)

		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error making user win")
			return
		}

		err = mines.AddBoxesWithMines(activeGame, NextBet.Box_No, multiplier)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error adding mines to boxes")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, "user won ", nil)
	}

}
