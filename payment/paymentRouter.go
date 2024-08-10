package payment

import (
	"strive_go/payment/controllers"
	"strive_go/payment/middlewares"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(superRoute *gin.RouterGroup) {
	booksRouter := superRoute.Group("/payment").Use(middlewares.CORSMiddleware())
	{
		booksRouter.POST("/addcash", controllers.AddCashController)

		booksRouter.POST("/createorder", controllers.CreateOrder)

		booksRouter.GET("/getuser", controllers.GetUserAccount)

	}
}
