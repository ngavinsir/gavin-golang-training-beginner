package postgres_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ngavinsir/golangtraining"
	"github.com/ngavinsir/golangtraining/internal/postgres"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type paymentCodesTestSuite struct {
	postgres.Suite
}

func TestSuitePaymentCodes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test Suite for PaymentCodes Repository")
	}

	dsn := os.Getenv("PG_TEST_DSN")

	if dsn == "" {
		dsn = "user=postgres password=password dbname=testing host=localhost port=54320 sslmode=disable"
	}

	paymentCodesSuite := &paymentCodesTestSuite{
		postgres.Suite{
			DSN:                     dsn,
			MigrationLocationFolder: "migrations",
		},
	}

	suite.Run(t, paymentCodesSuite)
}

func (s paymentCodesTestSuite) BeforeTest(suiteName, testName string) {
	ok, err := s.Migration.Up()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s paymentCodesTestSuite) AfterTest(suiteName, testName string) {
	ok, err := s.Migration.Down()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s paymentCodesTestSuite) TestCreatePaymentCodes() {
	uuid, err := uuid.NewRandom()
	s.Require().NoError(err)

	mockPaymentCodes := golangtraining.PaymentCode{
		ID:          uuid.String(),
		PaymentCode: "hello",
		Name:        "world",
	}

	testCases := []struct {
		desc           string
		repo           *postgres.PaymentCodesRepository
		expectedResult error
		ctx            context.Context
		reqBody        *golangtraining.PaymentCode
	}{
		{
			desc: "insert-success",
			repo: func() *postgres.PaymentCodesRepository {
				repo := postgres.NewPaymentCodesRepository(s.DBConn)
				return repo
			}(),
			expectedResult: nil,
			ctx:            context.TODO(),
			reqBody:        &mockPaymentCodes,
		},
		{
			desc: "context timeout. Too long to execute and already pass the limit context from parent",
			repo: func() *postgres.PaymentCodesRepository {
				repo := postgres.NewPaymentCodesRepository(s.DBConn)
				return repo
			}(),
			expectedResult: context.DeadlineExceeded,
			ctx: func() context.Context {
				bCtx := context.TODO()
				// Context already expired for 1 hour
				ctx, cancel := context.WithDeadline(bCtx, time.Now().Add(-1*time.Hour))
				defer cancel()
				return ctx
			}(),
			reqBody: &mockPaymentCodes,
		},
		{
			desc: "context Canceled by the caller",
			repo: func() *postgres.PaymentCodesRepository {
				repo := postgres.NewPaymentCodesRepository(s.DBConn)
				return repo
			}(),
			expectedResult: context.Canceled,
			ctx: func() context.Context {
				bCtx := context.TODO()
				// Context expired in 1 hour
				ctx, cancel := context.WithDeadline(bCtx, time.Now().Add(1*time.Hour))
				// Directly call cancel function
				defer cancel()
				return ctx
			}(),
			reqBody: &mockPaymentCodes,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			err := tC.repo.Create(tC.ctx, tC.reqBody)

			s.Require().Equal(tC.expectedResult, errors.Cause(err))
		})
	}
}
