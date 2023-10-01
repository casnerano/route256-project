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

func (rep *cartRepository) Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint32) (*model.Item, error) {
	row := rep.pgxpool.QueryRow(ctx, "INSERT INTO cart (user_id, sku_id, count) VALUES ($1, $2, $3) RETURNING id, user_id, sku_id, count", userID, sku, count)
	var item model.Item
	err := row.Scan(
		&item.ID,
		&item.UserID,
		&item.SKU,
		&item.Count,
	)
	return &item, err
}

func (rep *cartRepository) FindByUser(ctx context.Context, userID model.UserID) ([]*model.Item, error) {
	rows, err := rep.pgxpool.Query(ctx, "SELECT id, user_id, sku_id, count FROM cart where user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*model.Item

	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.ID,
			&item.UserID,
			&item.SKU,
			&item.Count,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (rep *cartRepository) DeleteBySKU(ctx context.Context, userID model.UserID, sku model.SKU) error {
	_, err := rep.pgxpool.Exec(ctx, "DELETE FROM cart WHERE user_id = $1 and sku_id = $2", userID, sku)
	return err
}

func (rep *cartRepository) DeleteByUser(ctx context.Context, userID model.UserID) error {
	_, err := rep.pgxpool.Exec(ctx, "DELETE FROM cart WHERE user_id = $1", userID)
	return err
}
