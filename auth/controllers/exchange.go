package controllers

import (
	"net/http"
	"strings"
	googleauth "strive_go/auth/services/googleAuth"
	striveauth "strive_go/auth/services/striveAuth"
	"strive_go/auth/utils"
	"strive_go/auth/utils/validators"
	"strive_go/db/functions"

	"github.com/gin-gonic/gin"
)

// ExchangeCode
// ExchangeCode                godoc
// @Title 		 GetGoogleAuthUrl
// @Summary		 Return Google auth url for google login of user
// @Description  Return Google auth url for google login of user
// @Tags         Auth
// @Param        code query string true "Code"
// @Produce      json
// @Success       200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/auth/exchange [put]
func ExchangeCode(c *gin.Context) {
	code := c.Query("code")

	authToken, err := googleauth.Exchange(code, c)
	if err != nil {
		utils.SendApiError(c, http.StatusBadRequest, "Invalid code")
		return
	}

	// TODO: handle dp oerations
	accessToken := authToken.AccessToken

	userInfo, err := googleauth.GetUserInfo(accessToken)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to get user Info from Google")
		return
	}

	email, ok := userInfo["email"].(string)
	if !ok {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to get user Info from Google")
		return
	}

	valid, err := validators.IsValidEmail(email)

	emailParts := strings.Split(email, "@")
	if len(emailParts) < 2 {
		utils.SendApiError(c, http.StatusInternalServerError, "Invalid Email format")
		return
	}

	username := emailParts[0]

	valid = validators.IsValidUsername(username)
	if !valid {
		utils.SendApiError(c, http.StatusInternalServerError, "Bad Username while using google auth")
		return
	}

	//username validation
	validUsername := validators.IsValidUsername(username)
	if !validUsername {
		utils.SendApiError(c, http.StatusInternalServerError, "Bad Username")
		return
	}
	userAuth, err := functions.InsertUserByGoogleAuth(username, email, authToken.RefreshToken)

	if err != nil {
		if err.Error() == "User already exists" {
			emailParts := strings.Split(email, "@")
			if len(emailParts) < 2 {
				utils.SendApiError(c, http.StatusInternalServerError, "Invalid Email format")
				return
			}

			username = emailParts[0]
		} else {
			utils.SendApiError(c, http.StatusInternalServerError, "Failed to register user")
			return
		}
	}

	//generate access token and refresh token

	accessToken, err = striveauth.GenerateAccessToken(userAuth.ID, userAuth.Email, userAuth.Username)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "Failed to generate acccess token")
		return
	}

	refreshToken, err := striveauth.GenerateRefreshToken(userAuth.ID)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, "failed to generate refresh token")
	}

	response := gin.H{
		"message":       "User logged in successfully",
		"username":      username,
		"email":         email,
		"phone":         nil,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	utils.SendSuccessResponse(c, http.StatusOK, "User logged in successfully", response)
}
