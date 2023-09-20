package order

import (
	"context"
	"errors"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
)

var (
	ErrNotFound        = errors.New("order not found")
	ErrCancelPaidOrder = errors.New("failed cancel paid order")
	ErrShipReserve     = errors.New("failed ship reserve")
	ErrCancelReserve   = errors.New("failed cancel reserve")
	ErrAddReserve      = errors.New("failed add reserve")
)

type order struct {
	repOrder repository.Order
	repStock repository.Stock
}

func New(repOrder repository.Order, repStock repository.Stock) *order {
	return &order{repOrder: repOrder, repStock: repStock}
}

func (o *order) Create(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error) {
	for _, item := range items {
		if err := o.repStock.AddReserve(ctx, item.SKU, item.Count); err != nil {
			return nil, ErrAddReserve
		}
	}

	createdOrder, err := o.repOrder.Add(ctx, userID, items)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return createdOrder, nil
}

func (o *order) GetInfo(ctx context.Context, orderID model.OrderID) (*model.Order, error) {
	foundOrder, err := o.repOrder.FindByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return foundOrder, nil
}

func (o *order) Payment(ctx context.Context, orderID model.OrderID) error {
	foundOrder, err := o.repOrder.FindByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}

		return err
	}

	for _, foundOrderItem := range foundOrder.Items {
		if err = o.repStock.ShipReserve(ctx, foundOrderItem.SKU, foundOrderItem.Count); err != nil {
			return ErrShipReserve
		}
	}

	return o.repOrder.ChangeStatus(ctx, orderID, model.OrderStatusPayed)
}

func (o *order) Cancel(ctx context.Context, orderID model.OrderID) error {
	foundOrder, err := o.repOrder.FindByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}

		return err
	}

	if foundOrder.Status == model.OrderStatusPayed {
		return ErrCancelPaidOrder
	}

	for _, foundOrderItem := range foundOrder.Items {
		if err = o.repStock.CancelReserve(ctx, foundOrderItem.SKU, foundOrderItem.Count); err != nil {
			return ErrCancelReserve
		}
	}

	return o.repOrder.ChangeStatus(ctx, orderID, model.OrderStatusCanceled)
}
