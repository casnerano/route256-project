package order

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"route256/loms/internal/model"
)

type createRequest struct {
	User  model.UserID `json:"user"`
	Items []item       `json:"items"`
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

	if createRequestStruct.User == 0 || len(createRequestStruct.Items) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	orderItems := make([]*model.OrderItem, len(createRequestStruct.Items))
	for _, rItem := range createRequestStruct.Items {
		orderItems = append(orderItems, &model.OrderItem{
			SKU:   rItem.SKU,
			Count: rItem.Count,
		})
	}

	createdOrder, err := h.orderService.Create(
		ctx,
		createRequestStruct.User,
		orderItems,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(createResponse{OrderID: createdOrder.ID}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
