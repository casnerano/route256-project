package memstore

import (
	"context"
	"errors"
	"math/rand"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
	"sync"
)

type stockStorage = map[model.SKU]*model.Stock

type stockRepository struct {
	mu    *sync.Mutex
	store stockStorage
}

func NewStockRepository() repository.Stock {
	return &stockRepository{
		mu:    &sync.Mutex{},
		store: make(stockStorage),
	}
}

var _ repository.Stock = (*stockRepository)(nil)

func (rep *stockRepository) FindBySKU(_ context.Context, sku model.SKU) (*model.Stock, error) {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	rep.setFixtureIfNotExist(sku)

	if stock, ok := rep.store[sku]; ok {
		return stock, nil
	}

	return nil, repository.ErrNotFound
}

func (rep *stockRepository) AddReserve(_ context.Context, sku model.SKU, count uint16) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	rep.setFixtureIfNotExist(sku)

	if stock, ok := rep.store[sku]; ok {
		if stock.Available < count {
			return errors.New("small quantity of availability")
		}

		rep.store[sku].Reserved += count
		rep.store[sku].Available -= count

		return nil
	}

	return repository.ErrNotFound
}

func (rep *stockRepository) CancelReserve(_ context.Context, sku model.SKU, count uint16) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	rep.setFixtureIfNotExist(sku)

	if stock, ok := rep.store[sku]; ok {
		if stock.Reserved < count {
			return errors.New("small quantity of reserve")
		}

		rep.store[sku].Available += count
		rep.store[sku].Reserved -= count

		return nil
	}

	return repository.ErrNotFound
}

func (rep *stockRepository) ShipReserve(_ context.Context, sku model.SKU, count uint16) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	rep.setFixtureIfNotExist(sku)

	if stock, ok := rep.store[sku]; ok {
		if stock.Reserved < count {
			return errors.New("small quantity of reserve")
		}

		rep.store[sku].Reserved -= count

		return nil
	}

	return repository.ErrNotFound
}

func (rep *stockRepository) setFixtureIfNotExist(sku model.SKU) {
	if _, ok := rep.store[sku]; !ok {
		rep.store[sku] = &model.Stock{
			Available: uint16(rand.Uint32() % 100),
			Reserved:  0,
		}
	}
}
