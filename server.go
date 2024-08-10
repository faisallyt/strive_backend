package main

import (
	"strive_go/auth"
	"strive_go/bets"
	_ "strive_go/docs"
	"strive_go/games"
	"strive_go/payment"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Strive Backend API
// @version 1.0
// @description This is the documentation for backend APIs of Strive. You can explore the API endpoints here.

func initRouter() *gin.Engine {
	app := gin.Default()

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := app.Group("/api/v1")

	auth.AuthRoutes(router)
	bets.BetsRoutes(router)
	games.GamesRoutes(router)
	payment.PaymentRoutes(router)

	return app
}
