package payments

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ngavinsir/golangtraining"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=mocks/mock_inquiries_service.go -package=mocks . InquiriesService
type InquiriesService interface {
	GetByTransactionID(ctx context.Context, id string) (golangtraining.Inquiry, error)
}

//go:generate mockgen -destination=mocks/mock_payments_repo.go -package=mocks . PaymentsRepository
type PaymentsRepository interface {
	Create(ctx context.Context, p *golangtraining.Payment) error
}

//go:generate mockgen -destination=mocks/mock_paymentcodes_service.go -package=mocks . PaymentCodesService
type PaymentCodesService interface {
	GetByPaymentCode(ctx context.Context, paymentCode string) (golangtraining.PaymentCode, error)
}

type PaymentsService struct {
	inquiriesService    InquiriesService
	paymentsRepo        PaymentsRepository
	paymentCodesService PaymentCodesService
}

func NewService(inquiriesService InquiriesService, paymentsRepo PaymentsRepository, paymentCodesService PaymentCodesService) *PaymentsService {
	return &PaymentsService{
		inquiriesService:    inquiriesService,
		paymentsRepo:        paymentsRepo,
		paymentCodesService: paymentCodesService,
	}
}

func (s PaymentsService) Create(ctx context.Context, p *golangtraining.Payment) (golangtraining.PaymentCode, error) {
	var pc golangtraining.PaymentCode

	_, err := s.inquiriesService.GetByTransactionID(ctx, p.TransactionID)
	if err != nil {
		return pc, err
	}

	pc, err = s.paymentCodesService.GetByPaymentCode(ctx, p.PaymentCode)
	if err != nil {
		return pc, err
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		err = errors.Wrap(err, "can't generate the UUID")
		return pc, err
	}
	p.ID = newUUID.String()

	now := time.Now().UTC()
	p.CreatedAt = now
	p.UpdatedAt = now

	err = s.paymentsRepo.Create(ctx, p)
	if err != nil {
		return pc, err
	}

	return pc, nil
}
