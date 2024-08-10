package controllers

import (
	"net/http"
	"strive_go/auth/utils"
	"strive_go/auth/utils/validators"
	"strive_go/db"
	"strive_go/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ChangePassword Return the response of userData provided by client.
// ChangePassword Controller                godoc
// @title		  ChangePasswordController
// @Summary       request for ChangePassword
// @Description   Returns the response of statusCode ok and changes Password.
// @Tags          Auth
// @Produce       json
// @Param         Authorization header string true "Bearer token"
// @Param         username   formData  string  true  "Username"
// @Param         oldPassword formData  string  true  "Old PassWord"
// @Param         newPassword formData string true "New Password"
// @Success       200        {object}  map[string]interface{}  "ok"
// @Router        /api/v1/auth/changePassword [post]
func ChangePassword(c *gin.Context) {
	var passwordData struct {
		Username    string `form:"username" binding:"required"`
		OldPassword string `form:"oldPassword" binding:"required"`
		NewPassword string `form:"newPassword" binding:"required"`
	}

	if err := c.ShouldBind(&passwordData); err != nil {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	//username validation
	ValidUsername := validators.IsValidUsername(passwordData.Username)

	if !ValidUsername {
		utils.SendApiError(c, http.StatusBadRequest, "Bad Username")
		return
	}

	//password validation
	ValidPassword := validators.IsValidPassword(passwordData.NewPassword)

	if !ValidPassword {
		utils.SendApiError(c, http.StatusBadRequest, "Bad Password")
		return
	}

	if passwordData.OldPassword == passwordData.NewPassword {
		// fmt.Println(passwordData.NewPassword)
		utils.SendApiError(c, http.StatusBadRequest, "New Pasword can not be same as Old password")
		return
	}

	var userAuth models.UserAuth

	db.Instance.Where("username=?", passwordData.Username).First(&userAuth)

	err := bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(passwordData.OldPassword))
	if err != nil {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordData.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to update password")
		return
	}
	userAuth.Password = string(hashedPassword)

	updatedResult := db.Instance.Save(&userAuth)

	if updatedResult.Error != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "something went wrong while updating password")
	}

	utils.SendSuccessResponse(c, http.StatusOK, "Password changed Succcessfully", nil)

}
