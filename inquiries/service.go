package inquiries

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngavinsir/golangtraining"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=mocks/mock_inquiries_repo.go -package=mocks . Repository
type Repository interface {
	Create(ctx context.Context, p *golangtraining.Inquiry) error
}

//go:generate mockgen -destination=mocks/mock_paymentcodes_service.go -package=mocks . PaymentCodesService
type PaymentCodesService interface {
	GetByPaymentCode(ctx context.Context, paymentCode string) (golangtraining.PaymentCode, error)
}

type InquiriesService struct {
	repo                Repository
	paymentCodesService PaymentCodesService
}

func NewService(repo Repository, paymentCodesService PaymentCodesService) *InquiriesService {
	return &InquiriesService{
		repo:                repo,
		paymentCodesService: paymentCodesService,
	}
}

func (s InquiriesService) Create(ctx context.Context, i *golangtraining.Inquiry) (golangtraining.PaymentCode, error) {
	p, err := s.paymentCodesService.GetByPaymentCode(ctx, i.PaymentCode)
	if err != nil {
		return p, err
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return p, err
	}
	i.ID = newUUID.String()

	now := time.Now().UTC()
	i.CreatedAt = now
	i.UpdatedAt = now

	err = s.repo.Create(ctx, i)
	if err != nil {
		return p, err
	}

	return p, nil
}
