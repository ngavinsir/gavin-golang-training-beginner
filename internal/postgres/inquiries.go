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

func (r InquiriesRepository) GetByTransactionID(ctx context.Context, id string) (golangtraining.Inquiry, error) {
	var res golangtraining.Inquiry
	sqlStatement := `SELECT * FROM inquiries where transaction_id=$1 limit 1`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)
	if err := row.Scan(
		&res.ID, &res.PaymentCode, &res.TransactionID, &res.Amount, &res.CreatedAt, &res.UpdatedAt,
	); err != nil {
		return res, err
	}

	return res, nil
}
