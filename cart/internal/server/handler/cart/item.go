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

type Modifier interface {
    Add(ctx context.Context, userID model.UserID, sku uint32, count uint16) error
    Delete(ctx context.Context, userID model.UserID, sku uint32) error
}

type itemAddRequest struct {
    User  model.UserID `json:"user"`
    SKU   uint32       `json:"SKU"`
    Count uint16       `json:"count"`
}

type itemDeleteRequest struct {
    User model.UserID `json:"user"`
    SKU  uint32       `json:"SKU"`
}

type itemHandler struct {
    modifier Modifier
}

func NewItemHandler(modifier Modifier) *itemHandler {
    return &itemHandler{modifier: modifier}
}

func (i *itemHandler) Add(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

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

    if err := i.modifier.Add(
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

func (i *itemHandler) Delete(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

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

    if err := i.modifier.Delete(ctx, itemDeleteRequestStruct.User, itemDeleteRequestStruct.SKU); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
