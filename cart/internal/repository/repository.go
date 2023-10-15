package repository

//go:generate mockgen -destination=mock/repository.go -source=repository.go

import (
	"context"
	"errors"
	"route256/cart/internal/model"
)

var ErrNotFound = errors.New("not row found")

type Cart interface {
	Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint32) (*model.Item, error)
	FindByUser(ctx context.Context, userID model.UserID) ([]*model.Item, error)
	DeleteBySKU(ctx context.Context, userID model.UserID, sku model.SKU) error
	DeleteByUser(ctx context.Context, userID model.UserID) error
}
