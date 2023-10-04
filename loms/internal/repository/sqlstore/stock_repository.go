package memstore

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
)

type stockRepository struct {
	pgxpool *pgxpool.Pool
}

func NewStockRepository(pgxpool *pgxpool.Pool) repository.Stock {
	return &stockRepository{
		pgxpool: pgxpool,
	}
}

var _ repository.Stock = (*stockRepository)(nil)

func (rep *stockRepository) FindBySKU(_ context.Context, sku model.SKU) (*model.Stock, error) {
	return nil, repository.ErrNotFound
}

func (rep *stockRepository) AddReserve(_ context.Context, sku model.SKU, count uint32) error {
	return repository.ErrNotFound
}

func (rep *stockRepository) CancelReserve(_ context.Context, sku model.SKU, count uint32) error {
	return repository.ErrNotFound
}

func (rep *stockRepository) ShipReserve(_ context.Context, sku model.SKU, count uint32) error {
	return repository.ErrNotFound
}

func (rep *stockRepository) setFixtureIfNotExist(sku model.SKU) {

}
