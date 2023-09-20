package cart

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"route256/cart/internal/server/handler"
	"runtime/debug"
	"time"

	"route256/cart/internal/model"
	cartService "route256/cart/internal/service/cart"
)

type listRequest struct {
	User model.UserID `json:"user"`
}

func (l *listRequest) valid() bool {
	return l.User != 0
}

type listItemResponse struct {
	SKU   model.SKU `json:"sku"`
	Count uint16    `json:"count"`
	Name  string    `json:"name"`
	Price uint32    `json:"price"`
}

type listResponse struct {
	Items      []*listItemResponse `json:"items"`
	TotalPrice uint32              `json:"totalPrice"`
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	listRequestStruct := listRequest{}
	if err := json.NewDecoder(r.Body).Decode(&listRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !listRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Second)
	defer cancel()

	list, err := h.service.List(ctx, listRequestStruct.User)
	if err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			handler.WriteResponse(w, listResponse{})
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	var response listResponse

	for key := range list {
		item := listItemResponse{
			SKU:   list[key].SKU,
			Count: list[key].Count,
			Name:  list[key].Name,
			Price: list[key].Price,
		}

		response.Items = append(response.Items, &item)
		response.TotalPrice += item.Price * uint32(item.Count)
	}

	handler.WriteResponse(w, response)
}
