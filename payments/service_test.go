package payments_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ngavinsir/golangtraining"
	"github.com/ngavinsir/golangtraining/payments"
	"github.com/ngavinsir/golangtraining/payments/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreateInquiry(t *testing.T) {
	mockPaymentCode := golangtraining.PaymentCode{
		PaymentCode: "code",
	}
	mockInquiry := golangtraining.Inquiry{
		PaymentCode:   "code",
		TransactionID: "id",
	}

	type resType struct {
		Res golangtraining.PaymentCode
		Err error
	}

	testCases := []struct {
		desc                string
		paymentsRepo        *mocks.MockPaymentsRepository
		paymentCodesService *mocks.MockPaymentCodesService
		inquiriesService    *mocks.MockInquiriesService
		publisher           *mocks.MockPublisher
		ctxTimeout          time.Duration
		ctx                 context.Context
		expectedReturn      resType
	}{
		{
			desc: "create payment inquiry - success",
			paymentsRepo: func() *mocks.MockPaymentsRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentsRepository(ctrl)

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
			inquiriesService: func() *mocks.MockInquiriesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockInquiriesService(ctrl)

				m.
					EXPECT().
					GetByTransactionID(gomock.Any(), gomock.Any()).
					Return(mockInquiry, nil)

				return m
			}(),
			publisher: func() *mocks.MockPublisher {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPublisher(ctrl)

				m.
					EXPECT().
					Publish(gomock.Any()).
					Return(nil)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
			expectedReturn: resType{
				Res: mockPaymentCode,
				Err: nil,
			},
		},
		{
			desc: "create inquiry - return error from repository",
			paymentsRepo: func() *mocks.MockPaymentsRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentsRepository(ctrl)

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
			inquiriesService: func() *mocks.MockInquiriesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockInquiriesService(ctrl)

				m.
					EXPECT().
					GetByTransactionID(gomock.Any(), gomock.Any()).
					Return(mockInquiry, nil)

				return m
			}(),
			publisher: func() *mocks.MockPublisher {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPublisher(ctrl)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
			expectedReturn: resType{
				Res: golangtraining.PaymentCode{},
				Err: errors.New("Unknown Error"),
			},
		},
		{
			desc: "create payment codes - return error from paymentCodesService GetByPaymentCode",
			paymentsRepo: func() *mocks.MockPaymentsRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentsRepository(ctrl)

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
			inquiriesService: func() *mocks.MockInquiriesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockInquiriesService(ctrl)

				m.
					EXPECT().
					GetByTransactionID(gomock.Any(), gomock.Any()).
					Return(mockInquiry, nil)

				return m
			}(),
			publisher: func() *mocks.MockPublisher {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPublisher(ctrl)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
			expectedReturn: resType{
				Res: golangtraining.PaymentCode{},
				Err: errors.New("Unknown Error"),
			},
		},
		{
			desc: "create payment inquiry - return error from inquiriesService getByTransactionID",
			paymentsRepo: func() *mocks.MockPaymentsRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentsRepository(ctrl)

				return m
			}(),
			paymentCodesService: func() *mocks.MockPaymentCodesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentCodesService(ctrl)

				return m
			}(),
			inquiriesService: func() *mocks.MockInquiriesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockInquiriesService(ctrl)

				m.
					EXPECT().
					GetByTransactionID(gomock.Any(), gomock.Any()).
					Return(golangtraining.Inquiry{}, errors.New("Unknown Error"))

				return m
			}(),
			publisher: func() *mocks.MockPublisher {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPublisher(ctrl)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
			expectedReturn: resType{
				Res: golangtraining.PaymentCode{},
				Err: errors.New("Unknown Error"),
			},
		},
		{
			desc: "create payment inquiry - error from publisher",
			paymentsRepo: func() *mocks.MockPaymentsRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentsRepository(ctrl)

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
			inquiriesService: func() *mocks.MockInquiriesService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockInquiriesService(ctrl)

				m.
					EXPECT().
					GetByTransactionID(gomock.Any(), gomock.Any()).
					Return(mockInquiry, nil)

				return m
			}(),
			publisher: func() *mocks.MockPublisher {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPublisher(ctrl)

				m.
					EXPECT().
					Publish(gomock.Any()).
					Return(errors.New("Unknown Error"))

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
			expectedReturn: resType{
				Res: golangtraining.PaymentCode{},
				Err: errors.New("Unknown Error"),
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := payments.NewService(tC.inquiriesService, tC.paymentsRepo, tC.paymentCodesService, tC.publisher)
			p, err := service.Create(tC.ctx, &golangtraining.Payment{})

			if tC.expectedReturn.Err != nil {
				require.Equal(t, tC.expectedReturn.Err.Error(), errors.Cause(err).Error())
			} else {
				require.Equal(t, p, mockPaymentCode)
			}
		})
	}
}
