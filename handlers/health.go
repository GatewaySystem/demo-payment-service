package handlers

import (
	"net/http"
	"time"

	"github.com/GatewaySystem/demo-payment-service/config"
	"github.com/gin-gonic/gin"
)

var startedAt = time.Now()

func HealthCheck(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":         "healthy",
			"service":        cfg.ServiceName,
			"version":        cfg.Version,
			"uptime_seconds": int(time.Since(startedAt).Seconds()),
		})
	}
}
