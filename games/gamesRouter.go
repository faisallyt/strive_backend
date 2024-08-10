package games

import (
	"strive_go/auth/middlewares"
	"strive_go/games/controllers/mines"

	"github.com/gin-gonic/gin"
)

func GamesRoutes(superRoute *gin.RouterGroup) {
	gameRouter := superRoute.Group("/games")
	{
		// gameRouter.POST("/start-mines")
	}
	gameRouter.Use(middlewares.AuthMiddlware)
	{
		gameRouter.POST("/start-mines", mines.MinesStartController)
	}
}
