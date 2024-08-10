package controllers

import (
	"fmt"
	"net/http"
	striveauth "strive_go/auth/services/striveAuth"
	"strive_go/auth/utils"
	"strive_go/auth/utils/validators"
	"strive_go/db"
	"strive_go/db/functions"
	"strive_go/models"

	"github.com/gin-gonic/gin"
)

// LoginController Return the response of authenticated user.
// LoginController                godoc
// @Title 		loginController
// @Summary		request for authenticate users
// @Description  use for authenticate registered user with the help of email and password.
// @Tags         Auth
// @Produce      json
// @Param        username formData string  false  "registered username"
// @Param        email formData string  false  "registered email"
// @Param        phone formData string  false  "registered phone"
// @Param        password formData string  true  "password"
// @Success      200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/auth/login [post]
func LoginController(c *gin.Context) {
	var loginData struct {
		Email    string `form:"email"`
		Username string `form:"username"`
		Phone    string `form:"Phone"`
		Password string `form:"password"`
	}

	if err := c.ShouldBind(&loginData); err != nil {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	//username validation

	validUsername := validators.IsValidUsername(loginData.Username)

	if loginData.Username != "" && !validUsername {
		utils.SendApiError(c, http.StatusBadRequest, "Bad Username")
		return
	}

	//phone validation

	validPhone := validators.IsValidPhone(loginData.Phone)
	if loginData.Phone != "" && !validPhone {
		utils.SendApiError(c, http.StatusBadRequest, "Bad Phone")
		return
	}

	//password validation

	validPassword := validators.IsValidPassword(loginData.Password)
	if loginData.Password != "" && !validPassword {
		utils.SendApiError(c, http.StatusBadRequest, "Bad Password")
		return
	}

	fmt.Println("User data", loginData.Email, "-", loginData.Username, "-", loginData.Phone, "-", loginData.Password)

	//Validate login credentials

	err := functions.ValidateLogin(loginData.Email, loginData.Username, loginData.Phone, loginData.Password)

	if err != nil {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	var userAuth models.UserAuth

	if loginData.Email != "" {
		db.Instance.Where("email=?", loginData.Email).First(&userAuth)
	} else if loginData.Username != "" {
		db.Instance.Where("username=?", loginData.Username).First(&userAuth)
	} else if loginData.Phone != "" {
		db.Instance.Where("phone=?", loginData.Phone).First(&userAuth)
	}
	if !userAuth.Verified {
		utils.SendApiError(c, http.StatusUnauthorized, "Otp not verified")
	}
	//generate access token and refresh token

	accessToken, err := striveauth.GenerateAccessToken(userAuth.ID, userAuth.Email, userAuth.Password)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to generate acccess token")
		return
	}

	refreshToken, err := striveauth.GenerateRefreshToken(userAuth.ID)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "failed to generate refresh token")
	}

	err = functions.UpdateRefreshToken(userAuth.ID, refreshToken)
	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to update refresh token")
		return
	}

	// userData, err := functions.GetUserInfo(userAuth.Username)

	type UserInfo struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		DOB      string `json:"dob"`
	}

	userInfo := UserInfo{
		Username: userAuth.Username,
		Email:    userAuth.Email,
		Phone:    userAuth.Phone,
	}

	response := gin.H{
		"message":       "Login Successfull",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"userData":      userInfo,
	}

	utils.SendSuccessResponse(c, http.StatusOK, "login Successfull", response)

}
