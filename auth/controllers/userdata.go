package controllers

import (
	"fmt"
	"net/http"
	"strive_go/auth/utils"

	"github.com/gin-gonic/gin"
)

// UserController Return the response of details of logged in user.
// UserController                godoc
// @title		 UserController
// @Summary      request for userData
// @Description  Returns the response of logged in user, all about user such as name, username, email...
// @Tags         Auth
// @Produce      json
// @Param         Authorization header string true "Bearer token"
// @Success      200  {object}  string  "ok"
// @Router       /api/v1/auth/user [get]
func GetUser(c *gin.Context) {
	fmt.Println("userData Controller")
	utils.SendSuccessResponse(c, http.StatusOK, "userdata fetched successfully", "success")
}
