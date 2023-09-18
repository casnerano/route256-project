package repository

import (
    "context"
    "errors"
    "route256/cart/internal/model"
)

var ErrNotFound = errors.New("not found")

type Cart interface {
    Add(ctx context.Context, userID model.UserID, item *model.CartItem) error
    FindByUser(ctx context.Context, userID model.UserID) ([]*model.CartItem, error)
    DeleteBySKU(ctx context.Context, userID model.UserID, sku uint32) error
    DeleteByUser(ctx context.Context, userID model.UserID) error
}
