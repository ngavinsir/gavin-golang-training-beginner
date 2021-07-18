package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ngavinsir/golangtraining"
)

//go:generate mockgen -destination=mocks/mock_inquiries_service.go -package=mocks . InquiriesService
type InquiriesService interface {
	Create(ctx context.Context, i *golangtraining.Inquiry) (golangtraining.PaymentCode, error)
}

type inquiriesHandler struct {
	service InquiriesService
}

func InitInquiriesHandler(r *httprouter.Router, service InquiriesService) {
	h := inquiriesHandler{
		service: service,
	}

	r.HandlerFunc(http.MethodPost, "/inquiry", h.create)
}

func (h inquiriesHandler) create(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var inquiry golangtraining.Inquiry
	if err = json.Unmarshal(b, &inquiry); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = validator.Struct(inquiry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.service.Create(req.Context(), &inquiry) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serveJSON(w, p, http.StatusOK)
}
