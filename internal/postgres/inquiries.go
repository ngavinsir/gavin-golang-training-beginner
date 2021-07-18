package postgres

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/golangtraining"
)

type InquiriesRepository struct {
	DB *sql.DB
}

func NewInquiriesRepository(db *sql.DB) *InquiriesRepository {
	return &InquiriesRepository{
		DB: db,
	}
}

func (r InquiriesRepository) Create(ctx context.Context, i *golangtraining.Inquiry) error {
	sqlStatement := `
		INSERT INTO inquiries (id, payment_code, transaction_id, amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(
		ctx, sqlStatement, i.ID, i.PaymentCode, i.TransactionID, i.Amount, i.CreatedAt, i.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
