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

type infoRequest struct {
	OrderID model.OrderID `json:"orderID"`
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

	if infoRequestStruct.OrderID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	foundOrder, err := h.orderService.GetInfo(
		ctx,
		infoRequestStruct.OrderID,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := infoResponse{
		Status: foundOrder.Status,
		User:   foundOrder.User,
	}

	for _, rItem := range foundOrder.Items {
		response.Items = append(response.Items, item{
			SKU:   rItem.SKU,
			Count: rItem.Count,
		})
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
