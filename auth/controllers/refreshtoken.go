package controllers

import (
	"net/http"
	striveauth "strive_go/auth/services/striveAuth"
	"strive_go/auth/utils"
	"strive_go/db/functions"

	"github.com/gin-gonic/gin"
)

// TokenController Return the access token in case of expiry.
// TokenController                godoc
// @title		 TokenRefresher
// @Summary      request for access token
// @Description  Returns the access token in case of expiry with the help of refresh token.
// @Tags         Auth
// @Produce      json
// @Param        token formData string  true  "Refresh token"
// @Success      200  {object}  string  "ok"
// @Router       /api/v1/auth/refreshtoken [post]
func TokenController(c *gin.Context) {

	refreshtoken := c.PostForm("token")
	if refreshtoken == "" {
		utils.SendApiError(c, http.StatusBadRequest, "token is required")
		return
	}

	// check if the token is valid
	claims, err := striveauth.VerifyRefreshToken(refreshtoken)
	if err != nil {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	// get user auth details
	userAuth, err := functions.GetUserAuthFromID(claims.ID)
	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// generate new access token
	accessToken, err := striveauth.GenerateAccessToken(userAuth.ID, userAuth.Email, userAuth.Username)

	if err != nil {
		utils.SendApiError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// send the access token
	response := gin.H{
		"access_token": accessToken,
	}

	utils.SendSuccessResponse(c, http.StatusOK, "token refresh successful", response)

}
