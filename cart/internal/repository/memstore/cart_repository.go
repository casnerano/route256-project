package memstore

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
	"sync"
)

// cartStorage is the user's cart storage.
// For each USER a map is stored SKU => COUNT.
type cartStorage = map[model.UserID]map[model.SKU]uint32

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

func (rep *cartRepository) Add(_ context.Context, userID model.UserID, sku model.SKU, count uint32) (*model.Item, error) {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	if _, ok := rep.store[userID]; !ok {
		rep.store[userID] = make(map[model.SKU]uint32)
	}

	rep.store[userID][sku] = count
	rep.counter++

	return &model.Item{
		ID:     rep.counter,
		UserID: userID,
		SKU:    sku,
		Count:  count,
	}, nil
}

func (rep *cartRepository) FindByUser(_ context.Context, userID model.UserID) ([]*model.Item, error) {
	if userItems, ok := rep.store[userID]; ok {
		rep.mu.RLock()
		defer rep.mu.RUnlock()

		var items []*model.Item
		for sku := range userItems {
			items = append(items, &model.Item{
				SKU:   sku,
				Count: userItems[sku],
			})
		}

		return items, nil
	}
	return nil, repository.ErrNotFound
}

func (rep *cartRepository) DeleteBySKU(_ context.Context, userID model.UserID, sku model.SKU) error {
	if _, ok := rep.store[userID][sku]; ok {
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
