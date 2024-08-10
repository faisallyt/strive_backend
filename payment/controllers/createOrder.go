package controllers

import (
	"fmt"
	"net/http"
	"strive_go/payment/utils"

	"github.com/gin-gonic/gin"
)

// CreateOrder use to request for add to wallet
// CreateOrder                godoc
// @Summary      request for withrawling the amount
// @Description  Returns the details of the money withraw by users
// @Tags         payment
// @Produce      json
// @Param        username query string true "username"
// @Param        amount formData float64 true "amount to be withdraw"
// @Success      200  {object}  map[string]interface{}  "ok"
// @Router       /api/v1/payment/createorder [post]
func CreateOrder(c *gin.Context) {
	fmt.Println("createOrder Controller")
	utils.SendSuccessResponse(c, http.StatusOK, "created order successfully", "ok")
}
