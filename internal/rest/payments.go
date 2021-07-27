package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ngavinsir/golangtraining"
)

//go:generate mockgen -destination=mocks/mock_payments_service.go -package=mocks . PaymentsService
type PaymentsService interface {
	Create(ctx context.Context, i *golangtraining.Payment) (golangtraining.PaymentCode, error)
}

type paymentsHandler struct {
	service PaymentsService
}

func InitPaymentsHandler(r *httprouter.Router, service PaymentsService) {
	h := paymentsHandler{
		service: service,
	}

	r.HandlerFunc(http.MethodPost, "/payment", h.create)
}

func (h paymentsHandler) create(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var payment golangtraining.Payment
	if err = json.Unmarshal(b, &payment); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = validator.Struct(payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.service.Create(req.Context(), &payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serveJSON(w, p, http.StatusOK)
}
