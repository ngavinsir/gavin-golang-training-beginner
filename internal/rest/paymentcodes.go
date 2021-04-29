package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	golangtraining "github.com/ngavinsir/golangtraining"
)

type paymentCodesHandler struct {
	repo golangtraining.IPaymentCodesRepository
}

func InitPaymentCodesHandler(m *http.ServeMux, repo golangtraining.IPaymentCodesRepository) {
	h := paymentCodesHandler{
		repo: repo,
	}

	m.HandleFunc("/payment-codes", h.create)
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
