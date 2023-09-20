package cart

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/internal/server/handler"
	cartService "route256/cart/internal/service/cart"
	"runtime/debug"
	"time"
)

type checkoutRequest struct {
	User model.UserID `json:"user"`
}

func (c *checkoutRequest) valid() bool {
	return c.User != 0
}

type checkoutResponse struct {
	OrderID model.OrderID `json:"orderID"`
}

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	checkoutRequestStruct := checkoutRequest{}
	if err := json.NewDecoder(r.Body).Decode(&checkoutRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkoutRequestStruct.valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	var err error
	var response checkoutResponse

	response.OrderID, err = h.service.Checkout(ctx, checkoutRequestStruct.User)
	if err != nil {
		if errors.Is(err, cartService.ErrEmptyCart) {
			handler.WriteResponseError(w, 0, err.Error())
			return
		}

		handler.WriteInternalError(w, err)
		return
	}

	handler.WriteResponse(w, response)
}
