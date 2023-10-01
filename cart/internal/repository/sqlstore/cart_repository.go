package sqlstore

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
)

type cartRepository struct {
	pgxpool *pgxpool.Pool
}

func NewCartRepository(pgxpool *pgxpool.Pool) repository.Cart {
	return &cartRepository{
		pgxpool: pgxpool,
	}
}

var _ repository.Cart = (*cartRepository)(nil)

func (rep *cartRepository) Add(_ context.Context, userID model.UserID, sku model.SKU, count uint32) (*model.Item, error) {
	return nil, nil
}

func (rep *cartRepository) FindByUser(ctx context.Context, userID model.UserID) ([]*model.Item, error) {
	return nil, repository.ErrNotFound
}

func (rep *cartRepository) DeleteBySKU(ctx context.Context, userID model.UserID, sku model.SKU) error {
	return repository.ErrNotFound
}

func (rep *cartRepository) DeleteByUser(ctx context.Context, userID model.UserID) error {
	return repository.ErrNotFound
}
