package postgres

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/golangtraining"
)

type PaymentsRepository struct {
	DB *sql.DB
}

func NewPaymentsRepository(db *sql.DB) *PaymentsRepository {
	return &PaymentsRepository{
		DB: db,
	}
}

func (r PaymentsRepository) Create(ctx context.Context, p *golangtraining.Payment) error {
	sqlStatement := `
		INSERT INTO payments (id, name, payment_code, transaction_id, amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.ExecContext(
		ctx, sqlStatement, p.ID, p.Name, p.PaymentCode, p.TransactionID, p.Amount, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
