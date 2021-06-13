package rest_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/ngavinsir/golangtraining"
	"github.com/ngavinsir/golangtraining/internal/rest"
	"github.com/ngavinsir/golangtraining/internal/rest/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		desc               string
		service            *mocks.MockService
		URL                string
		method             string
		reqBody            io.Reader
		expectedStatusCode int
	}{
		{
			desc: "success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil)

				return m
			}(),
			URL:    "/payment-codes",
			method: http.MethodPost,
			reqBody: strings.NewReader(`
				{
					"payment_code": "hello",
					"name": "world" 
				}
			`),
			expectedStatusCode: http.StatusCreated,
		},
		{
			desc: "service error",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(errors.New("Database Error"))

				return m
			}(),
			URL:    "/payment-codes",
			method: http.MethodPost,
			reqBody: strings.NewReader(`
			{
				"payment_code": "hello",
				"name": "world" 
			}
			`),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentCodesHandler(r, tC.service)

			req := httptest.NewRequest(tC.method, tC.URL, tC.reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			require.Equal(t, tC.expectedStatusCode, rec.Code)
		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		desc               string
		service            *mocks.MockService
		URL                string
		method             string
		reqBody            io.Reader
		expectedStatusCode int
	}{
		{
			desc: "success",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					GetByID(gomock.Any(), gomock.Any()).
					Return(golangtraining.PaymentCode{}, nil)

				return m
			}(),
			URL:                "/payment-codes/7e8a17ba-3d1a-44d6-873e-e653f3888bf1",
			method:             http.MethodGet,
			reqBody:            nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			desc: "service error",
			service: func() *mocks.MockService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockService(ctrl)

				m.
					EXPECT().
					GetByID(gomock.Any(), gomock.Any()).
					Return(golangtraining.PaymentCode{}, errors.New("Database Error"))

				return m
			}(),
			URL:                "/payment-codes/7e8a17ba-3d1a-44d6-873e-e653f3888bf1",
			method:             http.MethodGet,
			reqBody:            nil,
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentCodesHandler(r, tC.service)

			req := httptest.NewRequest(tC.method, tC.URL, tC.reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			require.Equal(t, tC.expectedStatusCode, rec.Code)
		})
	}
}