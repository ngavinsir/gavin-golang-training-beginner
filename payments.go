package golangtraining

import (
	"time"
)

type Payment struct {
	ID            string    `json:"id"`
	PaymentCode   string    `json:"payment_code"`
	TransactionID string    `json:"transaction_id"`
	Name          string    `json:"name"`
	Amount        string    `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
