package middlewares

import (
	"net/http"
	"strings"
	googleauth "strive_go/auth/services/googleAuth"
	"strive_go/auth/utils"
	"strive_go/db"
	"strive_go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GoogleAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid authorization header")
		c.AbortWithStatus(401)
		return
	}

	//Expecting format of Bearer token

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid authorization")
		c.AbortWithStatus(401)
		return
	}

	tokenString := parts[1]

	userInformation, err := googleauth.GetUserInfo(tokenString)

	if err != nil {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid authorization")
		c.AbortWithStatus(401)
		return
	}

	userEmail, ok := userInformation["email"].(string)

	if !ok {
		utils.SendApiError(c, http.StatusUnauthorized, "Invalid email in token ")
		c.AbortWithStatus(401)
		return
	}

	var user models.UserAuth
	if err := db.Instance.Where("email=?", userEmail).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.SendApiError(c, http.StatusUnauthorized, "User not found")
		} else {
			utils.SendApiError(c, http.StatusInternalServerError, "Database error")
			c.AbortWithStatus(http.StatusInternalServerError)
			return

		}
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	username := user.Username

	c.Set("username", username)

	c.Next()

}
