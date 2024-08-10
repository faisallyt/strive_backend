package auth

import (
	"strive_go/auth/controllers"
	"strive_go/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(superRoute *gin.RouterGroup) {
	authRouter := superRoute.Group("/auth")
	{
		authRouter.GET("/user", controllers.GetUser)

		authRouter.POST("/login", controllers.LoginController)

		authRouter.PUT("/user", controllers.RegisterController)

		authRouter.POST("/otp", controllers.VerifyOtpController)

		authRouter.POST("/refreshtoken", controllers.TokenController)

		authRouter.GET("/GoogleAuthUrl", controllers.GoogleAuthUrl)

		authRouter.PUT("/exchange", controllers.ExchangeCode)

	}
	authRouter.Use(middlewares.AuthMiddlware)
	{
		authRouter.POST("/changePassword", controllers.ChangePassword)
		authRouter.POST("/change-username", controllers.ChangeUsernameController)
	}
}
