package sqlstore

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
	"time"
)

type orderStatusOutboxRepository struct {
	provider Provider
}

func NewOrderStatusOutboxRepository(provider Provider) repository.OrderStatusOutbox {
	return &orderStatusOutboxRepository{
		provider: provider,
	}
}

var _ repository.OrderStatusOutbox = (*orderStatusOutboxRepository)(nil)

func (rep *orderStatusOutboxRepository) Add(ctx context.Context, orderId model.OrderID, orderStatus model.OrderStatus) (*model.OrderStatusOutbox, error) {
	orderStatusOutbox := model.OrderStatusOutbox{
		OrderID:     orderId,
		OrderStatus: orderStatus,
		CreatedAt:   time.Now(),
	}

	row := rep.provider.Store(ctx).QueryRow(
		ctx,
		`INSERT INTO "order_status_outbox" (order_id, order_status, created_at) VALUES ($1, $2, $3) RETURNING id`,
		orderStatusOutbox.OrderID,
		orderStatusOutbox.OrderStatus,
		orderStatusOutbox.CreatedAt.UTC(),
	)

	err := row.Scan(
		&orderStatusOutbox.ID,
	)

	if err != nil {
		return nil, err
	}

	return &orderStatusOutbox, nil
}

func (rep *orderStatusOutboxRepository) MarkAsDelivery(ctx context.Context, id int) error {
	_, err := rep.provider.Store(ctx).Exec(ctx, `UPDATE "order_status_outbox" set is_delivery = true where id = $1`, id)
	return err
}

func (rep *orderStatusOutboxRepository) FindUndelivered(ctx context.Context) ([]*model.OrderStatusOutbox, error) {
	rows, err := rep.provider.Store(ctx).Query(
		ctx,
		`SELECT id, order_id, order_status, is_delivery, created_at FROM "order_status_outbox" where is_delivery = false`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordersStatusOutbox []*model.OrderStatusOutbox
	for rows.Next() {
		var orderStatusOutbox model.OrderStatusOutbox
		if err = rows.Scan(
			&orderStatusOutbox.ID,
			&orderStatusOutbox.OrderID,
			&orderStatusOutbox.OrderStatus,
			&orderStatusOutbox.IsDelivery,
			&orderStatusOutbox.CreatedAt,
		); err != nil {
			return nil, err
		}

		ordersStatusOutbox = append(ordersStatusOutbox, &orderStatusOutbox)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ordersStatusOutbox, nil
}
