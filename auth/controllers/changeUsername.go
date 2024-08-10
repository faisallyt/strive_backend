package controllers

import (
	"net/http"
	"strive_go/auth/utils"
	"strive_go/auth/utils/validators"
	"strive_go/db/functions"

	"github.com/gin-gonic/gin"
)

// ChangeUsernameController handles the username change requests.
// ChangeUsernameController                godoc
// @Title       ChangeUsernameController
// @Summary     Change the username of the authenticated user
// @Description Allows an authenticated user to change their username
// @Tags        User
// @Produce     json
// @Param       Authorization header string true "Bearer token"
// @Param       new_username  formData string true "New username"
// @Success     200 {object} map[string]interface{} "Username changed successfully"
// @Failure     400 {object} map[string]interface{} "New username is required"
// @Failure     401 {object} map[string]interface{} "Unauthorized"
// @Failure     500 {object} map[string]interface{} "Failed to change username"
// @Router      /api/v1/user/change-username [post]
func ChangeUsernameController(c *gin.Context) {
	newUsername := c.PostForm("new_username")
	if newUsername == "" {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid new username")
		c.AbortWithStatus(400)
		return
	}

	//username validation
	ValidUsername := validators.IsValidUsername(newUsername)
	if !ValidUsername {
		utils.SendApiError(c, http.StatusBadRequest, "Bad new username")
		return
	}

	username, exists := c.Get("username")

	if !exists {
		utils.SendApiError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if res, err := functions.UserExists("username", newUsername); res == true && err == nil {
		utils.SendApiError(c, http.StatusConflict, "Username already exists")
		return
	}

	err := functions.ChangeUsernameInDB(username.(string), newUsername)
	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Database error")
		return
	}

	utils.SendSuccessResponse(c, http.StatusOK, "username Change Succesfully", newUsername)
}
