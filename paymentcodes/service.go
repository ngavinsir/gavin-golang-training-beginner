package paymentcodes

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngavinsir/golangtraining"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=mocks/mock_paymentcodes_repo.go -package=mocks . Repository
type Repository interface {
	Create(ctx context.Context, p *golangtraining.PaymentCode) error
	GetByID(ctx context.Context, id string) (golangtraining.PaymentCode, error)
}

type PaymentCodesService struct {
	repo Repository
}

func NewService(repo Repository) *PaymentCodesService {
	return &PaymentCodesService{
		repo: repo,
	}
}

func (s PaymentCodesService) Create(ctx context.Context, p *golangtraining.PaymentCode) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return err
	}
	p.ID = newUUID.String()

	if p.ExpirationDate.IsZero() {
		p.ExpirationDate = time.Now().AddDate(51, 0, 0).UTC()
	}

	now := time.Now().UTC()
	p.CreatedAt = now
	p.UpdatedAt = now

	err = s.repo.Create(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (s PaymentCodesService) GetByID(ctx context.Context, id string) (golangtraining.PaymentCode, error) {
	res, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
