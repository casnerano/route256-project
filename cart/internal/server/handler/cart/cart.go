package cart

import (
	"context"
	"route256/cart/internal/model"
)

type Service interface {
	Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint16) error
	Delete(ctx context.Context, userID model.UserID, sku model.SKU) error
	List(ctx context.Context, userID model.UserID) ([]*model.CartItemDetail, error)
	Clear(ctx context.Context, userID model.UserID) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}
