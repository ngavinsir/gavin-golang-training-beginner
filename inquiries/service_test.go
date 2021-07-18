package inquiries_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ngavinsir/golangtraining"
	"github.com/ngavinsir/golangtraining/inquiries"
	"github.com/ngavinsir/golangtraining/inquiries/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreateInquiry(t *testing.T) {
	mockPaymentCode := golangtraining.PaymentCode{
		Name: "paymentcode",
	}

	testCases := []struct {
		desc                string
		repo                *mocks.MockRepository
		paymentCodesService *mocks.MockPaymentCodesService
		ctxTimeout          time.Duration
		ctx                 context.Context
		expectedReturn      error
	}{
		{
			desc: "create payment codes - success",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				m.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil)

				return m
			}(),
			paymentCodesService: func() *mocks.MockPaymentCodesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentCodesService(ctrl)

				m.
					EXPECT().
					GetByPaymentCode(gomock.Any(), gomock.Any()).
					Return(mockPaymentCode, nil)

				return m
			}(),
			ctxTimeout:     time.Second * 1,
			ctx:            context.TODO(),
			expectedReturn: nil,
		},
		{
			desc: "create payment codes - return error from repository",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				m.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(errors.New("Unknown Error"))

				return m
			}(),
			paymentCodesService: func() *mocks.MockPaymentCodesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentCodesService(ctrl)

				m.
					EXPECT().
					GetByPaymentCode(gomock.Any(), gomock.Any()).
					Return(mockPaymentCode, nil)

				return m
			}(),
			ctxTimeout:     time.Second * 1,
			ctx:            context.TODO(),
			expectedReturn: errors.New("Unknown Error"),
		},
		{
			desc: "create payment codes - return error from paymentCodesService GetByPaymentCode",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				return m
			}(),
			paymentCodesService: func() *mocks.MockPaymentCodesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentCodesService(ctrl)

				m.
					EXPECT().
					GetByPaymentCode(gomock.Any(), gomock.Any()).
					Return(golangtraining.PaymentCode{}, errors.New("Unknown Error"))

				return m
			}(),
			ctxTimeout:     time.Second * 1,
			ctx:            context.TODO(),
			expectedReturn: errors.New("Unknown Error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := inquiries.NewService(tC.repo, tC.paymentCodesService)
			p, err := service.Create(tC.ctx, &golangtraining.Inquiry{})
			if err != nil {
				require.Equal(t, tC.expectedReturn.Error(), errors.Cause(err).Error())
			} else {
				require.Equal(t, p, mockPaymentCode)
			}
		})
	}
}
