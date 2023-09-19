package memstore

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
	"sync"
)

type orderStorage = map[model.SKU]*model.Stock

type orderRepository struct {
	mu    *sync.RWMutex
	store orderStorage
}

func NewOrderRepository() repository.Order {
	return &orderRepository{
		mu:    &sync.RWMutex{},
		store: make(orderStorage),
	}
}

var _ repository.Order = (*orderRepository)(nil)

func (rep *orderRepository) Add(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error) {
	return nil, nil
}

func (rep *orderRepository) FindByID(ctx context.Context, orderID model.OrderID) (*model.Order, error) {
	return nil, nil
}

func (rep *orderRepository) ChangeStatus(ctx context.Context, orderID model.OrderID, status model.OrderStatus) error {
	return nil
}
