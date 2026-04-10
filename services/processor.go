package services

import (
	"math/rand"
	"time"

	"github.com/GatewaySystem/demo-payment-service/models"
	"github.com/GatewaySystem/demo-payment-service/store"
	"github.com/google/uuid"
)

// FailureRate is the mock payment failure percentage (5% default).
var FailureRate = 5

type Metrics struct {
	SuccessCount    int64
	FailureCount    int64
	RefundCount     int64
	TotalProcessed  float64
	LastDurationMs  int64
}

var Stats = &Metrics{}

func ProcessPayment(req models.CreatePaymentRequest) *models.Payment {
	start := time.Now()

	// Simulate processing delay 200-500ms
	delay := time.Duration(200+rand.Intn(300)) * time.Millisecond
	time.Sleep(delay)

	currency := req.Currency
	if currency == "" {
		currency = "USD"
	}
	method := req.Method
	if method == "" {
		method = "credit_card"
	}

	status := models.StatusSuccess
	if rand.Intn(100) < FailureRate {
		status = models.StatusFailed
	}

	payment := &models.Payment{
		ID:        "pay_" + uuid.New().String()[:8],
		OrderID:   req.OrderID,
		Amount:    req.Amount,
		Currency:  currency,
		Method:    method,
		Status:    status,
		CreatedAt: time.Now(),
	}

	store.Get().Save(payment)

	Stats.LastDurationMs = time.Since(start).Milliseconds()
	if status == models.StatusSuccess {
		Stats.SuccessCount++
		Stats.TotalProcessed += req.Amount
	} else {
		Stats.FailureCount++
	}

	return payment
}

func RefundPayment(id string) *models.Payment {
	p := store.Get().Find(id)
	if p == nil {
		return nil
	}
	if p.Status != models.StatusSuccess {
		return nil
	}
	p.Status = models.StatusRefunded
	store.Get().Save(p)
	Stats.RefundCount++
	return p
}
