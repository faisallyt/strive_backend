package controllers

import (
	"fmt"
	"net/http"
	"strive_go/payment/utils"

	"github.com/gin-gonic/gin"
)

// GetUserAccount use to fetch the logged in user account
// GetUserAccount                godoc
// @Summary      Get the result of logged in user account
// @Description  Returns the details of the user account
// @Tags         payment
// @Produce      json
// @Param        username query string true "username"
// @Success      200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/payment/getuser [get]
func GetUserAccount(c *gin.Context) {
	fmt.Println("GetUserAccount Controller")
	utils.SendSuccessResponse(c, http.StatusOK, "user account fetched successfully", "ok")
}
