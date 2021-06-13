package paymentcodes_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ngavinsir/golangtraining"
	"github.com/ngavinsir/golangtraining/paymentcodes"
	"github.com/ngavinsir/golangtraining/paymentcodes/mocks"
	"github.com/stretchr/testify/require"
)

func TestGetByIDPaymentCode(t *testing.T) {
	mockPaymentCode := golangtraining.PaymentCode{}

	type resType struct {
		Res golangtraining.PaymentCode
		Err error
	}

	testCases := []struct {
		desc           string
		repo           *mocks.MockRepository
		ctxTimeout     time.Duration
		ctx            context.Context
		expectedReturn resType
	}{
		{
			desc: "getByID payment codes - success",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				m.
					EXPECT().
					GetByID(gomock.Any(), gomock.Any()).
					Return(mockPaymentCode, nil)

				return m
			}(),
			ctxTimeout:     time.Second * 1,
			ctx:            context.TODO(),
			expectedReturn: resType{
				Res: mockPaymentCode,
			},
		},
		{
			desc: "getByID payment codes - return error from repository",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				m.
					EXPECT().
					GetByID(gomock.Any(), gomock.Any()).
					Return(mockPaymentCode, errors.New("Unknown Error"))

				return m
			}(),
			ctxTimeout:     time.Second * 1,
			ctx:            context.TODO(),
			expectedReturn: resType{
				Err: errors.New("Unknown Error"),
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := paymentcodes.NewService(tC.repo)
			res, err := service.GetByID(tC.ctx, "")

			require.Equal(t, tC.expectedReturn.Res, res)
			require.Equal(t, tC.expectedReturn.Err, err)
		})
	}
}
