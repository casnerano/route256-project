package stock

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"route256/loms/internal/model"
	"runtime/debug"
	"time"
)

type Service interface {
	GetAvailable(ctx context.Context, sku model.SKU) (uint16, error)
}

type stockInfoRequest struct {
	SKU model.SKU `json:"sku"`
}

type stockInfoResponse struct {
	Count uint16 `json:"count"`
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	stockInfoRequestStruct := stockInfoRequest{}
	if err := json.NewDecoder(r.Body).Decode(&stockInfoRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	var err error
	response := stockInfoResponse{}

	response.Count, err = h.service.GetAvailable(ctx, stockInfoRequestStruct.SKU)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
