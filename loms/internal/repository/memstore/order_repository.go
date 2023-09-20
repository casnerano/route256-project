package memstore

import (
	"context"
	"sync"
	"time"

	"route256/loms/internal/model"
	"route256/loms/internal/repository"
)

type orderStorage = map[model.OrderID]*model.Order

type orderRepository struct {
	mu      *sync.RWMutex
	store   orderStorage
	counter int64
}

func NewOrderRepository() repository.Order {
	return &orderRepository{
		mu:      &sync.RWMutex{},
		store:   make(orderStorage),
		counter: 1000,
	}
}

var _ repository.Order = (*orderRepository)(nil)

func (rep *orderRepository) Add(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error) {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	order := model.Order{
		ID:        rep.counter,
		Status:    model.OrderStatusNew,
		User:      userID,
		Items:     items,
		CreatedAt: time.Now().UTC(),
	}

	rep.store[rep.counter] = &order
	rep.counter += 100

	return &order, nil
}

func (rep *orderRepository) FindByID(_ context.Context, orderID model.OrderID) (*model.Order, error) {
	rep.mu.RLock()
	defer rep.mu.RUnlock()

	if order, ok := rep.store[orderID]; ok {
		return order, nil
	}

	return nil, repository.ErrRowNotFound
}

func (rep *orderRepository) ChangeStatus(ctx context.Context, orderID model.OrderID, status model.OrderStatus) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	if order, ok := rep.store[orderID]; ok {
		order.Status = status
		return nil
	}

	return repository.ErrRowNotFound
}
