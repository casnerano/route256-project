package repository

import (
	"context"
	"errors"
	"route256/loms/internal/model"
	"time"
)

var ErrNotFound = errors.New("not found")

type Order interface {
	Add(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error)
	FindByID(ctx context.Context, id model.OrderID) (*model.Order, error)
	ChangeStatus(ctx context.Context, id model.OrderID, status model.OrderStatus) error
	FindByUnpaidStatusWithDuration(ctx context.Context, duration time.Duration) ([]*model.Order, error)
}

type OrderStatusOutbox interface {
	Add(ctx context.Context, orderId model.OrderID, orderStatus model.OrderStatus) (*model.OrderStatusOutbox, error)
	MarkAsDelivery(ctx context.Context, id int) error
	FindUndelivered(ctx context.Context) ([]*model.OrderStatusOutbox, error)
}

type Stock interface {
	FindBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error)
	AddReserve(ctx context.Context, sku model.SKU, count uint32) error
	CancelReserve(ctx context.Context, sku model.SKU, count uint32) error
	ShipReserve(ctx context.Context, sku model.SKU, count uint32) error
}
