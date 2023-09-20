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

	cartService "route256/cart/internal/service/cart"

	"route256/cart/internal/model"
)

type clearRequest struct {
	User model.UserID `json:"user"`
}

func (c *clearRequest) valid() bool {
	return c.User != 0
}

func (h *Handler) Clear(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	clearRequestStruct := clearRequest{}
	if err := json.NewDecoder(r.Body).Decode(&clearRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !clearRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	if err := h.service.Clear(ctx, clearRequestStruct.User); err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
