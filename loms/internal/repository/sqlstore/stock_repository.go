package sqlstore

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"math/rand"
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

func (rep *stockRepository) FindBySKU(ctx context.Context, sku model.SKU) (*model.Stock, error) {
	rep.setFixtureIfNotExist(ctx, sku)

	row := rep.pgxpool.QueryRow(ctx, `SELECT * FROM "stock" where sku_id = $1`, sku)
	var stock model.Stock
	err := row.Scan(
		&stock.ID,
		&stock.SKU,
		&stock.Available,
		&stock.Reserved,
	)
	if err != nil {
		return nil, err
	}

	return &stock, nil
}

func (rep *stockRepository) AddReserve(ctx context.Context, sku model.SKU, count uint32) error {
	rep.setFixtureIfNotExist(ctx, sku)

	_, err := rep.pgxpool.Exec(
		ctx,
		`UPDATE "stock" set reserved = reserved + $2, available = available - $2  where sku_id = $1 and available >= $2`,
		sku,
		count,
	)
	return err
}

func (rep *stockRepository) CancelReserve(ctx context.Context, sku model.SKU, count uint32) error {
	_, err := rep.pgxpool.Exec(
		ctx,
		`UPDATE "stock" set reserved = reserved - $2, available = available + $2  where sku_id = $1 and reserved >= $2`,
		sku,
		count,
	)
	return err
}

func (rep *stockRepository) ShipReserve(ctx context.Context, sku model.SKU, count uint32) error {
	_, err := rep.pgxpool.Exec(
		ctx,
		`UPDATE "stock" set reserved = reserved - $2 where sku_id = $1 and reserved >= $2`,
		sku,
		count,
	)
	return err
}

func (rep *stockRepository) setFixtureIfNotExist(ctx context.Context, sku model.SKU) {
	var id int
	err := rep.pgxpool.QueryRow(ctx, `SELECT id FROM "stock" WHERE sku_id = $1`, sku).Scan(&id)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		_, _ = rep.pgxpool.Exec(
			ctx,
			`INSERT INTO "stock" (sku_id, available, reserved) VALUES ($1, $2, $3)`,
			sku,
			rand.Uint32()%100,
			0,
		)
	}
}
