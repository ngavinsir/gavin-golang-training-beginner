package jobs

import (
	"context"
)

type PaymentCodesService interface {
	Expire(ctx context.Context) error
}

type ExpirePaymentCodesJob struct {
	PaymentCodesService PaymentCodesService
}

func NewExpirePaymentCodesJob(service PaymentCodesService) *ExpirePaymentCodesJob {
	return &ExpirePaymentCodesJob{
		PaymentCodesService: service,
	}
}

func (j ExpirePaymentCodesJob) Work(ctx context.Context) error {
	err := j.PaymentCodesService.Expire(ctx)
	if err != nil {
		return err
	}

	return nil
}
