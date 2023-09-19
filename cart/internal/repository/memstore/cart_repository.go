package memstore

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
	"sync"
)

// cartStorage is the user's cart storage.
// For each USER a map is stored SKU => COUNT.
type cartStorage = map[model.UserID]map[model.SKU]uint16

type cartRepository struct {
	mu    *sync.RWMutex
	store cartStorage
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
		rep.store[userID] = make(map[model.SKU]uint16)
	}
	rep.store[userID][item.SKU] = item.Count
	return nil
}

func (rep *cartRepository) FindByUser(_ context.Context, userID model.UserID) ([]*model.CartItem, error) {
	if userItems, ok := rep.store[userID]; ok {
		rep.mu.RLock()
		defer rep.mu.RUnlock()

		var items []*model.CartItem
		for sku := range userItems {
			items = append(items, &model.CartItem{
				SKU:   sku,
				Count: userItems[sku],
			})
		}
		return items, nil
	}
	return nil, repository.ErrRowNotFound
}

func (rep *cartRepository) DeleteBySKU(_ context.Context, userID model.UserID, sku model.SKU) error {
	if _, ok := rep.store[userID][sku]; ok {
		rep.mu.Lock()
		defer rep.mu.Unlock()

		delete(rep.store[userID], sku)
		return nil
	}
	return repository.ErrRowNotFound
}

func (rep *cartRepository) DeleteByUser(_ context.Context, userID model.UserID) error {
	if _, ok := rep.store[userID]; !ok {
		rep.mu.Lock()
		defer rep.mu.Unlock()

		delete(rep.store, userID)
		return nil
	}
	return repository.ErrRowNotFound
}
