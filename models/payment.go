package models

import "time"

type PaymentStatus string

const (
	StatusPending   PaymentStatus = "pending"
	StatusSuccess   PaymentStatus = "success"
	StatusFailed    PaymentStatus = "failed"
	StatusRefunded  PaymentStatus = "refunded"
)

type Payment struct {
	ID        string        `json:"id"`
	OrderID   string        `json:"order_id"`
	Amount    float64       `json:"amount"`
	Currency  string        `json:"currency"`
	Method    string        `json:"method"`
	Status    PaymentStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
}

type CreatePaymentRequest struct {
	OrderID  string  `json:"order_id" binding:"required"`
	Amount   float64 `json:"amount" binding:"required,gt=0"`
	Currency string  `json:"currency"`
	Method   string  `json:"method"`
}

type RefundRequest struct {
	Reason string `json:"reason"`
}
