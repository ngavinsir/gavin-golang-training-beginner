package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ngavinsir/golangtraining"
	"github.com/pkg/errors"
)

type paymentCodesRepository struct {
	DB *sql.DB
}

func NewPaymentCodesRepository(db *sql.DB) golangtraining.IPaymentCodesRepository {
	return &paymentCodesRepository{
		DB: db,
	}
}

func (p paymentCodesRepository) Create(ctx context.Context, s *golangtraining.PaymentCode) (err error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return
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
	_, err = p.DB.Exec(sqlStatement, s.ID, s.PaymentCode, s.Name, s.Status, s.ExpirationDate, s.CreatedAt, s.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "cannot insert new payment code into DB")
		return
	}

	return
}

func (p paymentCodesRepository) GetByID(ctx context.Context, ID string) (res golangtraining.PaymentCode, err error) {
	sqlStatement := `SELECT * FROM payment_code where id=$1`
	row := p.DB.QueryRow(sqlStatement, ID)
	if err = row.Scan(
		&res.ID, &res.PaymentCode, &res.Name, &res.Status,
		&res.ExpirationDate, &res.CreatedAt, &res.UpdatedAt,
	); err != nil {
		err = errors.Wrap(err, "cannot get payment code by ID from DB")
		return
	}

	return
}
