package order

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	orderService "route256/loms/internal/service/order"
	"runtime/debug"
	"time"

	"route256/loms/internal/model"
)

type payRequest struct {
	OrderID model.OrderID `json:"orderID"`
}

type payResponse struct{}

func (h *Handler) Pay(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	payRequestStruct := payRequest{}
	if err := json.NewDecoder(r.Body).Decode(&payRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if payRequestStruct.OrderID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	err := h.service.Payment(ctx, payRequestStruct.OrderID)
	if err != nil {
		if errors.Is(err, orderService.ErrNotFound) || errors.Is(err, orderService.ErrShipReserve) {
			// ... error response
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := payResponse{}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
