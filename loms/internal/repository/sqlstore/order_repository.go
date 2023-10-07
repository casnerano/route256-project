package sqlstore

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
	"time"
)

type orderRepository struct {
	pgxpool *pgxpool.Pool
}

func NewOrderRepository(pgxpool *pgxpool.Pool) repository.Order {
	return &orderRepository{
		pgxpool: pgxpool,
	}
}

var _ repository.Order = (*orderRepository)(nil)

func (rep *orderRepository) Add(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error) {
	order := model.Order{
		Status: model.OrderStatusNew,
		User:   userID,
		Items:  items,
	}

	row := rep.pgxpool.QueryRow(
		ctx,
		`INSERT INTO "order" (user_id, status, items) VALUES ($1, $2, $3) RETURNING id, created_at`,
		order.User,
		order.Status,
		order.Items,
	)

	err := row.Scan(
		&order.ID,
		&order.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (rep *orderRepository) FindByID(ctx context.Context, id model.OrderID) (*model.Order, error) {
	order := model.Order{
		ID: id,
	}

	row := rep.pgxpool.QueryRow(ctx, `SELECT user_id, status, items, created_at FROM "order" where id = $1`, order.ID)

	err := row.Scan(
		&order.User,
		&order.Status,
		&order.Items,
		&order.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return &order, nil
}

func (rep *orderRepository) ChangeStatus(ctx context.Context, id model.OrderID, status model.OrderStatus) error {
	_, err := rep.pgxpool.Exec(ctx, `UPDATE "order" set status = $1 where id = $2`, status, id)
	return err
}

func (rep *orderRepository) FindByUnpaidStatusWithDuration(ctx context.Context, duration time.Duration) ([]*model.Order, error) {
	rows, err := rep.pgxpool.Query(
		ctx,
		`SELECT id, user_id, status, items, created_at FROM "order" where status = $1 and created_at > $2`,
		model.OrderStatusNew,
		time.Now().Add(duration),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		var order model.Order
		if err = rows.Scan(
			&order.ID,
			&order.User,
			&order.Status,
			&order.Items,
			&order.CreatedAt,
		); err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
