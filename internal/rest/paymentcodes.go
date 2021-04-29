package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ngavinsir/golangtraining"
)

type repository interface {
	Create(ctx context.Context, p *golangtraining.PaymentCode) error
	GetByID(ctx context.Context, ID string) (golangtraining.PaymentCode, error)
}

type paymentCodesHandler struct {
	repo repository
}

func InitPaymentCodesHandler(r *httprouter.Router, repo repository) {
	h := paymentCodesHandler{
		repo: repo,
	}

	r.HandlerFunc(http.MethodPost, "/payment-codes", h.create)
	r.HandlerFunc(http.MethodGet, "/payment-codes/:id", h.get)
}

func (h paymentCodesHandler) create(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var paymentCode golangtraining.PaymentCode
	if err = json.Unmarshal(b, &paymentCode); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = validator.Struct(paymentCode); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.repo.Create(req.Context(), &paymentCode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serveJSON(w, paymentCode, http.StatusCreated)
}

func (h paymentCodesHandler) get(w http.ResponseWriter, req *http.Request) {
	params := httprouter.ParamsFromContext(req.Context())
	paymentCodeID := params.ByName("id")

	paymentCode, err := h.repo.GetByID(req.Context(), paymentCodeID)
	if err != nil {
		serveJSON(w, "", http.StatusNotFound)
		return
	}

	serveJSON(w, paymentCode, http.StatusOK)
}
