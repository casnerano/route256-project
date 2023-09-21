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

type cancelRequest struct {
	OrderID model.OrderID `json:"orderID"`
}

func (c *cancelRequest) valid() bool {
	return c.OrderID != 0
}

type cancelResponse struct{}

func (h *Handler) Cancel(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	cancelRequestStruct := cancelRequest{}
	if err := json.NewDecoder(r.Body).Decode(&cancelRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !cancelRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	err := h.service.Cancel(ctx, cancelRequestStruct.OrderID)
	if err != nil {
		if errors.Is(err, orderService.ErrNotFound) ||
			errors.Is(err, orderService.ErrCancelPaidOrder) ||
			errors.Is(err, orderService.ErrCancelReserve) {

			handler.WriteResponseError(w, 0, err.Error())
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	handler.WriteResponse(w, cancelResponse{})
}
