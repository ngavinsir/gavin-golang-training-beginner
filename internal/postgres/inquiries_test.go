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

type inquiriesTestSuite struct {
	postgres.Suite
}

func TestSuiteInquiries(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip the Test Suite for Inquiries Repository")
	}

	dsn := os.Getenv("PG_TEST_DSN")

	if dsn == "" {
		dsn = "user=postgres password=password dbname=testing host=localhost port=54320 sslmode=disable"
	}

	inquiriesSuite := &inquiriesTestSuite{
		postgres.Suite{
			DSN:                     dsn,
			MigrationLocationFolder: "migrations",
		},
	}

	suite.Run(t, inquiriesSuite)
}

func (s inquiriesTestSuite) BeforeTest(suiteName, testName string) {
	ok, err := s.Migration.Up()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s inquiriesTestSuite) AfterTest(suiteName, testName string) {
	ok, err := s.Migration.Down()
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s inquiriesTestSuite) TestCreateInquiries() {
	uuid, err := uuid.NewRandom()
	s.Require().NoError(err)

	mockInquiry := golangtraining.Inquiry{
		ID:            uuid.String(),
		PaymentCode:   "payment_code",
		TransactionID: "transaction_id",
	}

	testCases := []struct {
		desc           string
		repo           *postgres.InquiriesRepository
		expectedResult error
		ctx            context.Context
		reqBody        *golangtraining.Inquiry
	}{
		{
			desc: "insert-success",
			repo: func() *postgres.InquiriesRepository {
				repo := postgres.NewInquiriesRepository(s.DBConn)
				return repo
			}(),
			expectedResult: nil,
			ctx:            context.TODO(),
			reqBody:        &mockInquiry,
		},
		{
			desc: "context timeout. Too long to execute and already pass the limit context from parent",
			repo: func() *postgres.InquiriesRepository {
				repo := postgres.NewInquiriesRepository(s.DBConn)
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
			reqBody: &mockInquiry,
		},
		{
			desc: "context Canceled by the caller",
			repo: func() *postgres.InquiriesRepository {
				repo := postgres.NewInquiriesRepository(s.DBConn)
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
			reqBody: &mockInquiry,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			err := tC.repo.Create(tC.ctx, tC.reqBody)

			s.Require().Equal(tC.expectedResult, errors.Cause(err))
		})
	}
}

func (s inquiriesTestSuite) TestGetByTransactionID() {
	seedData := golangtraining.Inquiry{
		ID:            "7e8a17ba-3d1a-44d6-873e-e653f3888bf1",
		TransactionID: "transaction",
	}
	repo := postgres.NewInquiriesRepository(s.DBConn)
	repo.Create(context.Background(), &seedData)

	type resType struct {
		Res golangtraining.Inquiry
		Err error
	}

	testCases := []struct {
		desc           string
		repo           *postgres.InquiriesRepository
		input          string
		expectedReturn resType
		ctx            context.Context
	}{
		{
			desc: "valid transaction id",
			repo: func() *postgres.InquiriesRepository {
				repo := postgres.NewInquiriesRepository(s.DBConn)
				return repo
			}(),
			input: "transaction",
			expectedReturn: resType{
				Res: seedData,
			},
			ctx: context.TODO(),
		},
		{
			desc: "invalid transaction id",
			repo: func() *postgres.InquiriesRepository {
				repo := postgres.NewInquiriesRepository(s.DBConn)
				return repo
			}(),
			input: "invalid transaction id",
			expectedReturn: resType{
				Err: errors.New("sql: no rows in result set"),
			},
			ctx: context.TODO(),
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			res, err := tC.repo.GetByTransactionID(context.Background(), tC.input)
			if err != nil {
				s.Require().Equal(tC.expectedReturn.Err.Error(), errors.Cause(err).Error())
			}

			s.Require().Equal(tC.expectedReturn.Res, res)
		})
	}
}
