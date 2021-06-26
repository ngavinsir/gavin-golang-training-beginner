package postgres

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/golangtraining"
)

type PaymentCodesRepository struct {
	DB *sql.DB
}

func NewPaymentCodesRepository(db *sql.DB) *PaymentCodesRepository {
	return &PaymentCodesRepository{
		DB: db,
	}
}

func (r PaymentCodesRepository) Create(ctx context.Context, p *golangtraining.PaymentCode) error {
	sqlStatement := `
		INSERT INTO payment_code (id, payment_code, name, status, expiration_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.ExecContext(
		ctx, sqlStatement, p.ID, p.PaymentCode, p.Name, p.Status,
		p.ExpirationDate, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r PaymentCodesRepository) GetByID(ctx context.Context, id string) (golangtraining.PaymentCode, error) {
	var paymentCode golangtraining.PaymentCode
	sqlStatement := `SELECT * FROM payment_code where id=$1`
	row := r.DB.QueryRowContext(ctx, sqlStatement, id)
	if err := row.Scan(
		&paymentCode.ID, &paymentCode.PaymentCode, &paymentCode.Name, &paymentCode.Status,
		&paymentCode.ExpirationDate, &paymentCode.CreatedAt, &paymentCode.UpdatedAt,
	); err != nil {
		return paymentCode, err
	}

	return paymentCode, nil
}
