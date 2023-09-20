package cart

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"route256/cart/internal/service/cart"

	"route256/cart/internal/model"
)

type clearRequest struct {
	User model.UserID `json:"user"`
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

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	if err := h.modifier.Clear(ctx, clearRequestStruct.User); err != nil {
		if err == cart.ErrItemNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
