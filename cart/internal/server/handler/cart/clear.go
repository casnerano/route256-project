package cart

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"route256/cart/internal/model"
)

type Cleaner interface {
	Clear(ctx context.Context, userID model.UserID) error
}

type clearRequest struct {
	User model.UserID `json:"user"`
}

type clearHandler struct {
	cleaner Cleaner
}

func NewClearHandler(cleaner Cleaner) *clearHandler {
	return &clearHandler{cleaner: cleaner}
}

func (c *clearHandler) Clear(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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

	if err := c.cleaner.Clear(ctx, clearRequestStruct.User); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
