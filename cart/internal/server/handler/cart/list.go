package cart

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"route256/cart/internal/model"
	"route256/cart/internal/service/cart"
)

type listRequest struct {
	User model.UserID `json:"user"`
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

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Second)
	defer cancel()

	list, err := h.modifier.List(ctx, listRequestStruct.User)
	if err != nil {
		if err == cart.ErrItemNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
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

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
