package handlers

import (
	"net/http"
	"time"

	"github.com/GatewaySystem/demo-payment-service/services"
	"github.com/gin-gonic/gin"
)

type metricPoint struct {
	Name      string            `json:"name"`
	Value     float64           `json:"value"`
	Timestamp string            `json:"timestamp"`
	Labels    map[string]string `json:"labels"`
	Type      string            `json:"type"`
}

func GetMetrics(c *gin.Context) {
	now := time.Now().UTC().Format(time.RFC3339)
	s := services.Stats

	points := []metricPoint{
		{Name: "payment_processed_total", Value: float64(s.SuccessCount), Timestamp: now, Labels: map[string]string{"status": "success"}, Type: "counter"},
		{Name: "payment_processed_total", Value: float64(s.FailureCount), Timestamp: now, Labels: map[string]string{"status": "failed"}, Type: "counter"},
		{Name: "payment_amount_usd", Value: s.TotalProcessed, Timestamp: now, Labels: map[string]string{"service": "payment-service"}, Type: "histogram"},
		{Name: "payment_processing_duration_ms", Value: float64(s.LastDurationMs), Timestamp: now, Labels: map[string]string{"service": "payment-service"}, Type: "histogram"},
		{Name: "payment_refund_total", Value: float64(s.RefundCount), Timestamp: now, Labels: map[string]string{"status": "success"}, Type: "counter"},
	}

	c.JSON(http.StatusOK, gin.H{"metrics": points})
}
