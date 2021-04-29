package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ngavinsir/golangtraining"
	"github.com/pkg/errors"
)

type PaymentCodesRepository struct {
	DB *sql.DB
}

func NewPaymentCodesRepository(db *sql.DB) *PaymentCodesRepository {
	return &PaymentCodesRepository{
		DB: db,
	}
}

func (p PaymentCodesRepository) Create(ctx context.Context, s *golangtraining.PaymentCode) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return err
	}
	s.ID = newUUID.String()
	now := time.Now().UTC()
	s.CreatedAt = now
	s.UpdatedAt = now

	if s.ExpirationDate.IsZero() {
		s.ExpirationDate = time.Now().AddDate(51, 0, 0).UTC()
	}

	sqlStatement := `
		INSERT INTO payment_code (id, payment_code, name, status, expiration_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = p.DB.ExecContext(
		ctx, sqlStatement, s.ID, s.PaymentCode, s.Name, s.Status, 
		s.ExpirationDate, s.CreatedAt, s.UpdatedAt,
	)
	if err != nil {
		err = fmt.Errorf("%s: %w","cannot insert new payment code into DB", err)
		return err
	}

	return nil
}

func (p PaymentCodesRepository) GetByID(ctx context.Context, ID string) (golangtraining.PaymentCode, error) {
	var paymentCode golangtraining.PaymentCode
	sqlStatement := `SELECT * FROM payment_code where id=$1`
	row := p.DB.QueryRowContext(ctx, sqlStatement, ID)
	if err := row.Scan(
		&paymentCode.ID, &paymentCode.PaymentCode, &paymentCode.Name, &paymentCode.Status,
		&paymentCode.ExpirationDate, &paymentCode.CreatedAt, &paymentCode.UpdatedAt,
	); err != nil {
		err = fmt.Errorf("%s: %w","cannot get payment code by ID from DB", err)
		return paymentCode, err
	}

	return paymentCode, nil
}
