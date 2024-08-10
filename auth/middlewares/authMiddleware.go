package middlewares

import (
	"net/http"
	"strings"
	striveauth "strive_go/auth/services/striveAuth"
	"strive_go/auth/utils"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your_secret_key")

func AuthMiddlware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid authorization header")
		c.AbortWithStatus(401)
		return
	}

	//Expecting format of Bearer token

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid authorization header")
		c.AbortWithStatus(401)
		return
	}

	tokeString := parts[1]

	claims, err := striveauth.VerifyAccessToken(tokeString)
	if err != nil {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid token")
		c.AbortWithStatus(401)
		return
	}

	username := claims.Username

	c.Set("username", username)

	c.Next()

}
