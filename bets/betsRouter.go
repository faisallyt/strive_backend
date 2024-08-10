package bets

import (
	"strive_go/bets/controllers"
	"strive_go/bets/middlewares"

	"github.com/gin-gonic/gin"
)

func BetsRoutes(superRoute *gin.RouterGroup) {
	booksRouter := superRoute.Group("/bet").Use(middlewares.CORSMiddleware())
	{
		booksRouter.GET("/getallbets", controllers.GetAllBets)

		booksRouter.GET("/getuserbets", controllers.GetUserBets)

		booksRouter.POST("/placebet", controllers.PlaceBetController)
	}
}
