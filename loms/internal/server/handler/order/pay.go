package order

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"route256/loms/internal/model"
	"route256/loms/internal/server/handler"
	orderService "route256/loms/internal/service/order"
	"runtime/debug"
	"time"
)

type payRequest struct {
	OrderID model.OrderID `json:"orderID"`
}

func (p *payRequest) valid() bool {
	return p.OrderID != 0
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

	if !payRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	err := h.service.Payment(ctx, payRequestStruct.OrderID)
	if err != nil {
		if errors.Is(err, orderService.ErrNotFound) || errors.Is(err, orderService.ErrShipReserve) {
			handler.WriteResponseError(w, 0, err.Error())
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	handler.WriteResponse(w, payResponse{})
}
