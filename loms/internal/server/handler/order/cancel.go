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

type cancelRequest struct {
	OrderID model.OrderID `json:"orderID"`
}

type cancelResponse struct{}

func (h *Handler) Cancel(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Failed close request body: %s\n", debug.Stack())
		}
	}()

	cancelRequestStruct := cancelRequest{}
	if err := json.NewDecoder(r.Body).Decode(&cancelRequestStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if cancelRequestStruct.OrderID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Millisecond)
	defer cancel()

	err := h.service.Cancel(ctx, cancelRequestStruct.OrderID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := cancelResponse{}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
