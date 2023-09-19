package repository

import (
	"context"
	"errors"
	"route256/loms/internal/model"
)

var ErrRowNotFound = errors.New("not row found")

type Order interface {
	Add(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error)
	FindByID(ctx context.Context, orderID model.OrderID) (*model.Order, error)
	ChangeStatus(ctx context.Context, orderID model.OrderID, status model.OrderStatus) error
}

type Stock interface {
	FindBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error)
	AddReserve(ctx context.Context, sku model.SKU, count uint64) error
	CancelReserve(ctx context.Context, sku model.SKU, count uint64) error
}
