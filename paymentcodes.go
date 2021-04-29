package golangtraining

import (
	"context"
	"time"
)

type PaymentCode struct {
	ID             string    `json:"id"`
	PaymentCode    string    `json:"payment_code"`
	Name           string    `json:"name" validate:"required"`
	Status         string    `json:"status"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type IPaymentCodesRepository interface {
	Create(ctx context.Context, p *PaymentCode) error
	GetByID(ctx context.Context, ID string) (PaymentCode, error)
}
