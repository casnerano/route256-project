package sqlstore

import (
	"context"
	"errors"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
	"time"

	"github.com/jackc/pgx/v5"
)

type orderRepository struct {
	provider Provider
}

func NewOrderRepository(provider Provider) repository.Order {
	return &orderRepository{
		provider: provider,
	}
}

var _ repository.Order = (*orderRepository)(nil)

func (rep *orderRepository) Add(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error) {
	order := model.Order{
		Status:    model.OrderStatusNew,
		User:      userID,
		Items:     items,
		CreatedAt: time.Now(),
	}

	row := rep.provider.Store(ctx).QueryRow(
		ctx,
		`INSERT INTO "order" (user_id, status, items, created_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		order.User,
		order.Status,
		order.Items,
		order.CreatedAt.UTC(),
	)

	err := row.Scan(
		&order.ID,
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

	row := rep.provider.Store(ctx).QueryRow(ctx, `SELECT user_id, status, items, created_at FROM "order" where id = $1`, order.ID)

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
	_, err := rep.provider.Store(ctx).Exec(ctx, `UPDATE "order" set status = $1 where id = $2`, status, id)
	return err
}

func (rep *orderRepository) FindByUnpaidStatusWithDuration(ctx context.Context, duration time.Duration) ([]*model.Order, error) {
	rows, err := rep.provider.Store(ctx).Query(
		ctx,
		`SELECT id, user_id, status, items, created_at FROM "order" where status = $1 and created_at < $2`,
		model.OrderStatusNew,
		time.Now().UTC().Add(-duration),
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
