package order

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"route256/loms/internal/server/handler"
	orderService "route256/loms/internal/service/order"
	"runtime/debug"
	"time"

	"route256/loms/internal/model"
)

type infoRequest struct {
	OrderID model.OrderID `json:"orderID"`
}

func (i *infoRequest) valid() bool {
	return i.OrderID != 0
}

type infoResponse struct {
	Status model.OrderStatus `json:"status"`
	User   model.UserID      `json:"user"`
	Items  []item            `json:"items"`
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	infoRequestStruct := infoRequest{}
	if err := json.NewDecoder(r.Body).Decode(&infoRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !infoRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	foundOrder, err := h.service.GetInfo(ctx, infoRequestStruct.OrderID)
	if err != nil {
		if errors.Is(err, orderService.ErrNotFound) {
			handler.WriteResponseError(w, 0, err.Error())
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	response := infoResponse{Status: foundOrder.Status, User: foundOrder.User}

	for _, rItem := range foundOrder.Items {
		response.Items = append(response.Items, item{
			SKU:   rItem.SKU,
			Count: rItem.Count,
		})
	}

	handler.WriteResponse(w, response)
}
