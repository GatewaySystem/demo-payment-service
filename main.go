package main

import (
	"fmt"
	"log"

	"github.com/GatewaySystem/demo-payment-service/config"
	"github.com/GatewaySystem/demo-payment-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health & metrics
	r.GET("/health", handlers.HealthCheck(cfg))
	r.GET("/metrics", handlers.GetMetrics)

	// Payment routes
	r.POST("/api/v1/payments", handlers.CreatePayment)
	r.GET("/api/v1/payments/:id", handlers.GetPayment)
	r.GET("/api/v1/payments", handlers.ListPayments)
	r.POST("/api/v1/payments/:id/refund", handlers.RefundPayment)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("%s running on port %s", cfg.ServiceName, cfg.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
