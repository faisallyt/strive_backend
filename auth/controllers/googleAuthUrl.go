package controllers

import (
	"net/http"
	googleauth "strive_go/auth/services/googleAuth"
	"strive_go/auth/utils"

	"github.com/gin-gonic/gin"
)

// GetGoogleAuthUrl Return Google auth url for google login of user
// GetGoogleAuthUrl                godoc
// @Title 		 GetGoogleAuthUrl
// @Summary		 Return Google auth url for google login of user
// @Description  Return Google auth url for google login of user
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/auth/GoogleAuthUrl [get]
func GoogleAuthUrl(c *gin.Context) {
	url, state := googleauth.GetURL()
	utils.SendSuccessResponse(c, http.StatusOK, "url generated", gin.H{"url": url, "state": state})
}
