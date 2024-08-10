package controllers

import (
	"net/http"
	striveauth "strive_go/auth/services/striveAuth"
	"strive_go/auth/utils"
	"strive_go/db"
	"strive_go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// VerifyController  Returns the succesful or failure otp verification after matching.
// VerifyController                godoc
// @title		 OTP verification
// @Summary      request and verify otp for users
// @Description  Returns the succesful or failure otp verification after matching.
// @Tags         Auth
// @Produce      json
// @Param         username formData string true "Username"
// @Param         otp         formData string true "OTP verification"
// @Success       200         {object} map[string]interface{} "ok"
// @Router       /api/v1/auth/otp [post]
func VerifyOtpController(c *gin.Context) {
	var otpData struct {
		Username string `form:"username" binding:"required"`
		Otp      uint   `form:"otp" binding:"required"`
	}

	if err := c.ShouldBind(&otpData); err != nil {
		utils.SendApiError(c, http.StatusBadRequest, "invalid request data")
		return
	}

	if otpData.Otp < 100000 || otpData.Otp > 999999 {
		utils.SendApiError(c, http.StatusBadRequest, "Bad OTP")
		return
	}

	var userAuth models.UserAuth

	result := db.Instance.Where("username=?", otpData.Username).First(&userAuth)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.SendApiError(c, http.StatusNotFound, "User not found")
		} else {
			utils.SendApiError(c, http.StatusInternalServerError, "Database error")
		}
		return
	}

	if userAuth.OTP != otpData.Otp {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid otp")
		return
	}
	userAuth.Verified = true
	userAuth.OTP = 0

	if err := db.Instance.Save(&userAuth).Error; err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	//generate access token and refresh token

	accessToken, err := striveauth.GenerateAccessToken(userAuth.ID, userAuth.Email, userAuth.Username)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to generate acccess token")
		return
	}

	refreshToken, err := striveauth.GenerateRefreshToken(userAuth.ID)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "failed to generate refresh token")
	}

	response := gin.H{
		"message":       "User updated successfully",
		"username":      userAuth.Username,
		"email":         userAuth.Email,
		"phone":         userAuth.Phone,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	utils.SendSuccessResponse(c, http.StatusOK, "verification succesfull, User Registered", response)
}
