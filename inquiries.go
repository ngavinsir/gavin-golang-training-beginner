package golangtraining

import (
	"time"
)

type Inquiry struct {
	ID            string    `json:"id"`
	PaymentCode   string    `json:"payment_code" validate:"required"`
	TransactionID string    `json:"transaction_id" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
