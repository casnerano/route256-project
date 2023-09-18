package cart

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"route256/cart/internal/model"
	"runtime/debug"
	"time"
)

type itemAddRequest struct {
	User  model.UserID `json:"user"`
	SKU   model.SKU    `json:"SKU"`
	Count uint16       `json:"count"`
}

type itemDeleteRequest struct {
	User model.UserID `json:"user"`
	SKU  model.SKU    `json:"SKU"`
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

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	if err := h.modifier.Add(
		ctx,
		itemAddRequestStruct.User,
		itemAddRequestStruct.SKU,
		itemAddRequestStruct.Count,
	); err != nil {
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

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	if err := h.modifier.Delete(ctx, itemDeleteRequestStruct.User, itemDeleteRequestStruct.SKU); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
