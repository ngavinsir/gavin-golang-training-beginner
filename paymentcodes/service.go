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
	Expire(ctx context.Context) error
}

//go:generate mockgen -destination=mocks/mock_users.go -package=mocks . Users
type Users interface {
	GetUsers(ctx context.Context) (golangtraining.User, error)
}

type PaymentCodesService struct {
	repo  Repository
	users Users
}

func NewService(repo Repository, users Users) *PaymentCodesService {
	return &PaymentCodesService{
		repo:  repo,
		users: users,
	}
}

func (s PaymentCodesService) Create(ctx context.Context, p *golangtraining.PaymentCode) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return err
	}
	p.ID = newUUID.String()
	p.Status = "ACTIVE"

	if p.ExpirationDate.IsZero() {
		p.ExpirationDate = time.Now().AddDate(51, 0, 0).UTC()
	}

	now := time.Now().UTC()
	p.CreatedAt = now
	p.UpdatedAt = now

	user, err := s.users.GetUsers(ctx)
	if err != nil {
		return err
	}
	p.Name += user.Name

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

func (s PaymentCodesService) Expire(ctx context.Context) error {
	err := s.repo.Expire(ctx)
	if err != nil {
		return err
	}

	return nil
}
