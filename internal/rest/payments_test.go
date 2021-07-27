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

func TestCreatePayment(t *testing.T) {
	testCases := []struct {
		desc               string
		service            *mocks.MockPaymentsService
		URL                string
		method             string
		reqBody            io.Reader
		expectedStatusCode int
	}{
		{
			desc: "success",
			service: func() *mocks.MockPaymentsService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentsService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(golangtraining.PaymentCode{}, nil)

				return m
			}(),
			URL:    "/payment",
			method: http.MethodPost,
			reqBody: strings.NewReader(`
			{
				"transaction_id": "123",
				"amount ": 12345, 
				"name": "lechsaLeanne Graham", 
				"payment_code": "abc123"
			}
			`),
			expectedStatusCode: http.StatusOK,
		},
		{
			desc: "service error",
			service: func() *mocks.MockPaymentsService {
				ctrl := gomock.NewController(t)
				m := mocks.NewMockPaymentsService(ctrl)

				m.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(golangtraining.PaymentCode{}, errors.New("Database Error"))

				return m
			}(),
			URL:    "/payment",
			method: http.MethodPost,
			reqBody: strings.NewReader(`
			{
				"transaction_id": "123",
				"amount ": 12345, 
				"name": "lechsaLeanne Graham", 
				"payment_code": "abc123"
			}
		`),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := httprouter.New()
			rest.InitPaymentsHandler(r, tC.service)

			req := httptest.NewRequest(tC.method, tC.URL, tC.reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)
			require.Equal(t, tC.expectedStatusCode, rec.Code)
		})
	}
}
