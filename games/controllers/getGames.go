package controllers

import (
	"fmt"
	"net/http"
	"strive_go/games/utils"

	"github.com/gin-gonic/gin"
)

// GetGames use to fetch the all games available
// GetGames                godoc
// @Summary      Get the result of available games
// @Description  Returns the list of all the game available on site
// @Tags         Dice
// @Produce      json
// @Success      200  {object}  string  "ok"
// @Router       /api/v1/games/getgames [get]
func GetGames(c *gin.Context) {
	fmt.Println("getGames Controller")
	utils.SendSuccessResponse(c, http.StatusOK, "controller respone", "ok")
}
