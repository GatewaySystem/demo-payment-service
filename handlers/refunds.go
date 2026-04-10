package handlers

import (
	"net/http"

	"github.com/GatewaySystem/demo-payment-service/services"
	"github.com/gin-gonic/gin"
)

func RefundPayment(c *gin.Context) {
	id := c.Param("id")
	payment := services.RefundPayment(id)
	if payment == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment not found or not eligible for refund"})
		return
	}
	c.JSON(http.StatusOK, payment)
}
