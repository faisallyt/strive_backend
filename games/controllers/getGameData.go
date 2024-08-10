package controllers

import (
	"fmt"
	"net/http"
	"strive_go/games/utils"

	"github.com/gin-gonic/gin"
)

// GetGameData use to fetch the data related to the game
// GetGameData                godoc
// @Summary      Get the result of gamedata
// @Description  Returns the analytics data of particular game
// @Tags         Dice
// @Produce      json
// @Param        gameid query string true "gameid"
// @Success      200  {object}  string  "ok"
// @Router       /api/v1/games/getgamedata [get]
func GetGameData(c *gin.Context) {
	fmt.Println("gamedata Controller")
	utils.SendSuccessResponse(c, http.StatusOK, "gamedata successfully", "ok")
}
