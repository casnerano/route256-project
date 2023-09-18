package memstore

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
	"sync"
)

type cartStorage = map[model.UserID]map[uint32]*model.CartItem

type cartRepository struct {
	mu      *sync.RWMutex
	store   cartStorage
	counter int
}

func NewCartRepository() repository.Cart {
	return &cartRepository{
		mu:    &sync.RWMutex{},
		store: make(cartStorage),
	}
}

var _ repository.Cart = (*cartRepository)(nil)

func (rep *cartRepository) Add(_ context.Context, userID model.UserID, item *model.CartItem) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()
	if _, ok := rep.store[userID]; !ok {
		rep.store[userID] = make(map[uint32]*model.CartItem)
	}
	rep.store[userID][item.SKU] = item
	return nil
}

func (rep *cartRepository) FindByUser(_ context.Context, userID model.UserID) ([]*model.CartItem, error) {
	if userItems, ok := rep.store[userID]; ok {
		rep.mu.RLock()
		defer rep.mu.RUnlock()

		var items []*model.CartItem
		for key := range userItems {
			items = append(items, userItems[key])
		}
		return items, nil
	}
	return nil, repository.ErrNotFound
}

func (rep *cartRepository) DeleteBySKU(_ context.Context, userID model.UserID, sku uint32) error {
	if _, ok := rep.store[userID][sku]; !ok {
		rep.mu.Lock()
		defer rep.mu.Unlock()

		delete(rep.store[userID], sku)
		return nil
	}
	return repository.ErrNotFound
}

func (rep *cartRepository) DeleteByUser(_ context.Context, userID model.UserID) error {
	if _, ok := rep.store[userID]; !ok {
		rep.mu.Lock()
		defer rep.mu.Unlock()

		delete(rep.store, userID)
		return nil
	}
	return repository.ErrNotFound
}
