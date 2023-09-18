package cart

import (
	"context"
	"route256/cart/internal/model"
)

type Modifier interface {
	Add(ctx context.Context, userID model.UserID, sku uint32, count uint16) error
	Delete(ctx context.Context, userID model.UserID, sku uint32) error
	List(ctx context.Context, userID model.UserID) ([]*model.CartItem, error)
	Clear(ctx context.Context, userID model.UserID) error
}

type Handler struct {
	modifier Modifier
}

func NewHandler(modifier Modifier) *Handler {
	return &Handler{modifier: modifier}
}
