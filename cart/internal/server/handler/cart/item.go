package cart

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"route256/cart/internal/model"
	"route256/cart/internal/service/cart"
)

type itemAddRequest struct {
	User  model.UserID `json:"user"`
	SKU   model.SKU    `json:"sku"`
	Count uint16       `json:"count"`
}

type itemDeleteRequest struct {
	User model.UserID `json:"user"`
	SKU  model.SKU    `json:"sku"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	itemAddRequestStruct := itemAddRequest{}
	if err := json.NewDecoder(r.Body).Decode(&itemAddRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if itemAddRequestStruct.User == 0 || itemAddRequestStruct.SKU == 0 || itemAddRequestStruct.Count == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	if err := h.modifier.Add(
		ctx,
		itemAddRequestStruct.User,
		itemAddRequestStruct.SKU,
		itemAddRequestStruct.Count,
	); err != nil {
		if errors.Is(err, cart.ErrPIMProductNotFound) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	itemDeleteRequestStruct := itemDeleteRequest{}
	if err := json.NewDecoder(r.Body).Decode(&itemDeleteRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if itemDeleteRequestStruct.User == 0 || itemDeleteRequestStruct.SKU == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	if err := h.modifier.Delete(ctx, itemDeleteRequestStruct.User, itemDeleteRequestStruct.SKU); err != nil {
		if errors.Is(err, cart.ErrItemNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
