package order

import (
	"context"

	"route256/loms/internal/model"
)

type item struct {
	SKU   model.SKU `json:"sku"`
	Count uint16    `json:"count"`
}

type OrderService interface {
	Create(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error)
	GetInfo(ctx context.Context, orderID model.OrderID) (*model.Order, error)
	Payment(ctx context.Context, orderID model.OrderID) error
	Cancel(ctx context.Context, orderID model.OrderID) error
}

type Handler struct {
	orderService OrderService
}

func NewHandler(orderService OrderService) *Handler {
	return &Handler{orderService: orderService}
}
