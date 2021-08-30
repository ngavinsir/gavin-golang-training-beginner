package postgres

import (
	"context"
	"database/sql"
	"time"

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

func (r PaymentCodesRepository) GetByPaymentCode(ctx context.Context, paymentCode string) (golangtraining.PaymentCode, error) {
	var res golangtraining.PaymentCode
	sqlStatement := `SELECT * FROM payment_code where payment_code=$1 limit 1`
	row := r.DB.QueryRowContext(ctx, sqlStatement, paymentCode)
	if err := row.Scan(
		&res.ID, &res.PaymentCode, &res.Name, &res.Status,
		&res.ExpirationDate, &res.CreatedAt, &res.UpdatedAt,
	); err != nil {
		return res, err
	}

	return res, nil
}

func (r PaymentCodesRepository) Expire(ctx context.Context) error {
	updatedAt := time.Now().UTC()

	sqlStatement := `UPDATE payment_code SET updated_at=$1, status='INACTIVE' WHERE status = 'ACTIVE' and expiration_date <= $2`
	_, err := r.DB.ExecContext(ctx, sqlStatement, updatedAt, time.Now().UTC())
	if err != nil {
		return err
	}

	return nil
}
