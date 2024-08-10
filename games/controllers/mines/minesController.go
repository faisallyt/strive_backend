package mines

import (
	"net/http"
	"strive_go/auth/utils"
	"strive_go/db/functions/mines"

	"github.com/gin-gonic/gin"
)

type StartGameRequest struct {
	Amount     float64 `json:"amount" binding:"required" description:"amount"`
	MinesCount int64   `json:"mines_count" binding:"required" `
}

func MinesStartController(c *gin.Context) {

	username, exists := c.Get("username")

	if !exists {
		utils.SendApiError(c, http.StatusUnauthorized, "User is not authorized ")
		return
	}
	var request StartGameRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	isActive, gameId, err := mines.CheckActiveGame(username.(string))

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Error checking active games")
		return
	}

	if isActive {
		if gameId == nil {
			utils.SendApiError(c, http.StatusBadRequest, "Already you have an active game")
			return
		}

		//return the game that is currently active
		currentGame, err := mines.FetchCurrentGame(gameId)
		if err != nil {
			utils.SendApiError(c, http.StatusInternalServerError, "Error fetching current game")
			return
		}
		utils.SendSuccessResponse(c, http.StatusOK, "There is an Active game ", currentGame)
	}

	// if no active games

	err = mines.StartNewGame(username.(string), request.Amount, request.MinesCount)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "error starting a new Game")
	}

	utils.SendSuccessResponse(c, http.StatusOK, "New Game Started Successfully", nil)

}
