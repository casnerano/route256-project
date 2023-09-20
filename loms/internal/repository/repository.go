package repository

import (
	"context"
	"errors"

	"route256/loms/internal/model"
)

var ErrNotFound = errors.New("not found")

type Order interface {
	Add(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error)
	FindByID(ctx context.Context, orderID model.OrderID) (*model.Order, error)
	ChangeStatus(ctx context.Context, orderID model.OrderID, status model.OrderStatus) error
}

type Stock interface {
	FindBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error)
	AddReserve(ctx context.Context, sku model.SKU, count uint16) error
	CancelReserve(ctx context.Context, sku model.SKU, count uint16) error
	ShipReserve(ctx context.Context, sku model.SKU, count uint16) error
}
