package controllers

import (
	"fmt"
	"net/http"
	"strive_go/bets/utils"

	"github.com/gin-gonic/gin"
)

// GetUserBets Return the response of authenticated user.
// GetUserBets                godoc
// @Title 		GetUserBets Controller
// @Summary		request for users all bets
// @Description  use for fetch the users all bets and detailed information
// @Tags         Bet
// @Produce      json
// @Param        username query string  true  "registered username"
// @Success      200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/bet/getuserbets [get]
func GetUserBets(c *gin.Context) {
	fmt.Println(" UserBets controller")
	utils.SendSuccessResponse(c, http.StatusOK, "user bets successfully fetched", "ok")
}
