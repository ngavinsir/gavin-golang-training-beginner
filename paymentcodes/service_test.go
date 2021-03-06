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

func TestGetByID(t *testing.T) {
	mockPaymentCode := golangtraining.PaymentCode{}

	type resType struct {
		Res golangtraining.PaymentCode
		Err error
	}

	testCases := []struct {
		desc           string
		repo           *mocks.MockRepository
		users          *mocks.MockUsers
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
			users: func() *mocks.MockUsers {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockUsers(ctrl)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
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
			users: func() *mocks.MockUsers {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockUsers(ctrl)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
			expectedReturn: resType{
				Err: errors.New("Unknown Error"),
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := paymentcodes.NewService(tC.repo, tC.users)
			res, err := service.GetByID(tC.ctx, "")

			require.Equal(t, tC.expectedReturn.Res, res)
			require.Equal(t, tC.expectedReturn.Err, err)
		})
	}
}

func TestGetByPaymentCode(t *testing.T) {
	mockPaymentCode := golangtraining.PaymentCode{}

	type resType struct {
		Res golangtraining.PaymentCode
		Err error
	}

	testCases := []struct {
		desc           string
		repo           *mocks.MockRepository
		users          *mocks.MockUsers
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
					GetByPaymentCode(gomock.Any(), gomock.Any()).
					Return(mockPaymentCode, nil)

				return m
			}(),
			users: func() *mocks.MockUsers {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockUsers(ctrl)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
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
					GetByPaymentCode(gomock.Any(), gomock.Any()).
					Return(mockPaymentCode, errors.New("Unknown Error"))

				return m
			}(),
			users: func() *mocks.MockUsers {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockUsers(ctrl)

				return m
			}(),
			ctxTimeout: time.Second * 1,
			ctx:        context.TODO(),
			expectedReturn: resType{
				Err: errors.New("Unknown Error"),
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := paymentcodes.NewService(tC.repo, tC.users)
			res, err := service.GetByPaymentCode(tC.ctx, "")

			require.Equal(t, tC.expectedReturn.Res, res)
			require.Equal(t, tC.expectedReturn.Err, err)
		})
	}
}

func TestCreatePaymentCode(t *testing.T) {
	testCases := []struct {
		desc           string
		repo           *mocks.MockRepository
		users          *mocks.MockUsers
		ctxTimeout     time.Duration
		ctx            context.Context
		expectedReturn error
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
			users: func() *mocks.MockUsers {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockUsers(ctrl)

				m.
					EXPECT().
					GetUsers(gomock.Any()).
					Return(golangtraining.User{
						ID:   1,
						Name: "Leanne Graham",
					}, nil)

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
			users: func() *mocks.MockUsers {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockUsers(ctrl)

				m.
					EXPECT().
					GetUsers(gomock.Any()).
					Return(golangtraining.User{
						ID:   1,
						Name: "Leanne Graham",
					}, nil)

				return m
			}(),
			ctxTimeout:     time.Second * 1,
			ctx:            context.TODO(),
			expectedReturn: errors.New("Unknown Error"),
		},
		{
			desc: "create payment codes - return error from users",
			repo: func() *mocks.MockRepository {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockRepository(ctrl)

				return m
			}(),
			users: func() *mocks.MockUsers {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockUsers(ctrl)

				m.
					EXPECT().
					GetUsers(gomock.Any()).
					Return(golangtraining.User{}, errors.New("Unknown Error"))

				return m
			}(),
			ctxTimeout:     time.Second * 1,
			ctx:            context.TODO(),
			expectedReturn: errors.New("Unknown Error"),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := paymentcodes.NewService(tC.repo, tC.users)
			err := service.Create(tC.ctx, &golangtraining.PaymentCode{})

			require.Equal(t, tC.expectedReturn, err)
		})
	}
}
