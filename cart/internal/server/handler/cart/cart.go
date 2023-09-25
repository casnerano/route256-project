package cart

import (
	"context"
	"route256/cart/internal/model"
	pb "route256/cart/pkg/proto/cart/v1"
)

type Service interface {
	Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint32) error
	Delete(ctx context.Context, userID model.UserID, sku model.SKU) error
	List(ctx context.Context, userID model.UserID) ([]*model.ItemDetail, error)
	Clear(ctx context.Context, userID model.UserID) error
	Checkout(ctx context.Context, userID model.UserID) (model.OrderID, error)
}

type Handler struct {
	pb.UnimplementedCartServer
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}
