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

type createRequest struct {
	User  model.UserID `json:"user"`
	Items []item       `json:"items"`
}

func (c *createRequest) valid() bool {
	return c.User != 0 && len(c.Items) != 0
}

type createResponse struct {
	OrderID model.OrderID
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	createRequestStruct := createRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !createRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	orderItems := make([]*model.OrderItem, 0, len(createRequestStruct.Items))
	for _, rItem := range createRequestStruct.Items {
		orderItems = append(orderItems, &model.OrderItem{
			SKU:   rItem.SKU,
			Count: rItem.Count,
		})
	}

	createdOrder, err := h.service.Create(ctx, createRequestStruct.User, orderItems)
	if err != nil {
		if errors.Is(err, orderService.ErrAddReserve) {
			handler.WriteResponseError(w, 0, err.Error())
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	handler.WriteResponse(w, createResponse{OrderID: createdOrder.ID})
}
