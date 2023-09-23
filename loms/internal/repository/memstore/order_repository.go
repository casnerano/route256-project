package memstore

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
	"sync"
	"time"
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

	return nil, repository.ErrNotFound
}

func (rep *orderRepository) ChangeStatus(_ context.Context, orderID model.OrderID, status model.OrderStatus) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	if order, ok := rep.store[orderID]; ok {
		order.Status = status
		return nil
	}

	return repository.ErrNotFound
}

func (rep *orderRepository) FindByUnpaidStatusWithDuration(_ context.Context, duration time.Duration) ([]*model.Order, error) {
	rep.mu.RLock()
	defer rep.mu.RUnlock()

	orders := make([]*model.Order, 0)
	for _, order := range rep.store {
		if order.Status == model.OrderStatusNew && order.CreatedAt.Add(duration).Before(time.Now()) {
			orders = append(orders, order)
		}
	}

	return orders, nil
}