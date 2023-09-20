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
	"route256/cart/internal/server/handler"
	cartService "route256/cart/internal/service/cart"
)

type itemAddRequest struct {
	User  model.UserID `json:"user"`
	SKU   model.SKU    `json:"sku"`
	Count uint16       `json:"count"`
}

func (i *itemAddRequest) valid() bool {
	return i.User != 0 && i.SKU != 0 && i.Count != 0
}

type itemDeleteRequest struct {
	User model.UserID `json:"user"`
	SKU  model.SKU    `json:"sku"`
}

func (i *itemDeleteRequest) valid() bool {
	return i.User != 0 && i.SKU != 0
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

	if !itemAddRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	if err := h.service.Add(
		ctx,
		itemAddRequestStruct.User,
		itemAddRequestStruct.SKU,
		itemAddRequestStruct.Count,
	); err != nil {
		if errors.Is(err, cartService.ErrPIMProductNotFound) || errors.Is(err, cartService.ErrPIMLowAvailability) {
			handler.WriteResponseError(w, 0, err.Error())
			return
		}

		handler.WriteInternalError(w, err)
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

	if !itemDeleteRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	if err := h.service.Delete(ctx, itemDeleteRequestStruct.User, itemDeleteRequestStruct.SKU); err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
