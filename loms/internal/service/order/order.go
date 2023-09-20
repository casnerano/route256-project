package order

import (
	"context"
	"errors"
	"fmt"

	"route256/loms/internal/model"
	"route256/loms/internal/repository"
)

var ErrCancelPaidOrder = errors.New("failed cancel paid order")

type order struct {
	repOrder repository.Order
	repStock repository.Stock
}

func New(repOrder repository.Order, repStock repository.Stock) *order {
	return &order{repOrder: repOrder, repStock: repStock}
}

func (o *order) Create(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error) {
	return o.repOrder.Add(ctx, userID, items)
}

func (o *order) GetInfo(ctx context.Context, orderID model.OrderID) (*model.Order, error) {
	return o.repOrder.FindByID(ctx, orderID)
}

// TODO: context processing
func (o *order) Payment(ctx context.Context, orderID model.OrderID) error {
	fOrder, err := o.repOrder.FindByID(ctx, orderID)
	if err != nil {
		return err

	}

	for _, oItem := range fOrder.Items {
		if err = o.repStock.ShipReserve(ctx, oItem.SKU, uint64(oItem.Count)); err != nil {
			return fmt.Errorf("failed ship reserve: %w", err)
		}
	}

	o.repOrder.ChangeStatus(ctx, orderID, model.OrderStatusPayed)
	return nil
}

// TODO: context processing
func (o *order) Cancel(ctx context.Context, orderID model.OrderID) error {
	fOrder, err := o.repOrder.FindByID(ctx, orderID)
	if err != nil {
		return err

	}

	if fOrder.Status == model.OrderStatusPayed {
		return ErrCancelPaidOrder
	}

	for _, oItem := range fOrder.Items {
		if err = o.repStock.CancelReserve(ctx, oItem.SKU, uint64(oItem.Count)); err != nil {
			return fmt.Errorf("failed cencel reserve: %w", err)
		}
	}

	o.repOrder.ChangeStatus(ctx, orderID, model.OrderStatusCanceled)
	return nil
}
