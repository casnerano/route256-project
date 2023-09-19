package memstore

import (
	"context"
	"errors"
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

func (rep *stockRepository) FindBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error) {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	stock, ok := rep.store[sku]

	if !ok {
		stub := model.Stock{
			Available: 0,
			Reserved:  0,
		}
		rep.store[sku] = &stub
		return &stub, nil
	}

	return stock, nil
}

func (rep *stockRepository) AddReserve(ctx context.Context, sku model.SKU, count uint64) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	if stock, ok := rep.store[sku]; ok {
		if stock.Available < count {
			return errors.New("small quantity of availability")
		}

		rep.store[sku].Reserved += count
		rep.store[sku].Available -= count

		return nil
	}

	return repository.ErrRowNotFound
}

func (rep *stockRepository) CancelReserve(ctx context.Context, sku model.SKU, count uint64) error {
	rep.mu.Lock()
	defer rep.mu.Unlock()

	if stock, ok := rep.store[sku]; ok {
		if stock.Reserved < count {
			return errors.New("small quantity of reserve")
		}

		rep.store[sku].Available += count
		rep.store[sku].Reserved -= count

		return nil
	}

	return repository.ErrRowNotFound
}
