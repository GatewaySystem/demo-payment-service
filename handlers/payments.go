package handlers

import (
	"net/http"

	"github.com/GatewaySystem/demo-payment-service/models"
	"github.com/GatewaySystem/demo-payment-service/services"
	"github.com/GatewaySystem/demo-payment-service/store"
	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var req models.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := services.ProcessPayment(req)
	c.JSON(http.StatusCreated, payment)
}

func GetPayment(c *gin.Context) {
	id := c.Param("id")
	payment := store.Get().Find(id)
	if payment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func ListPayments(c *gin.Context) {
	payments := store.Get().List()
	c.JSON(http.StatusOK, gin.H{"payments": payments, "count": len(payments)})
}
