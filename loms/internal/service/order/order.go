package order

import (
	"context"
	"errors"
	"fmt"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
	"time"
)

var (
	ErrNotFound        = errors.New("order not found")
	ErrCancelPaidOrder = errors.New("failed cancel paid order")
	ErrShipReserve     = errors.New("failed ship reserve")
	ErrCancelReserve   = errors.New("failed cancel reserve")
	ErrAddReserve      = errors.New("failed add reserve")
	ErrPayedOrder      = errors.New("failed payed order")
)

type Order struct {
	statusSender StatusSender
	transactor   repository.Transactor
	repOrder     repository.Order
	repStock     repository.Stock
}

func New(statusSender StatusSender, transactor repository.Transactor, repOrder repository.Order, repStock repository.Stock) *Order {
	return &Order{
		statusSender: statusSender,
		transactor:   transactor,
		repOrder:     repOrder,
		repStock:     repStock,
	}
}

func (o *Order) Create(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error) {
	var createdOrder *model.Order

	err := o.transactor.RunRepeatableRead(ctx, func(txCtx context.Context) error {
		for _, item := range items {
			if err := o.repStock.AddReserve(txCtx, item.SKU, item.Count); err != nil {
				return ErrAddReserve
			}
		}

		var err error
		createdOrder, err = o.repOrder.Add(txCtx, userID, items)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				return ErrNotFound
			}

			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if sendErr := o.statusSender.Send(createdOrder.ID, createdOrder.Status); sendErr != nil {
		fmt.Println("Failed send create order status: ", sendErr.Error())
	}

	return createdOrder, nil
}

func (o *Order) GetInfo(ctx context.Context, orderID model.OrderID) (*model.Order, error) {
	foundOrder, err := o.repOrder.FindByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return foundOrder, nil
}

func (o *Order) Payment(ctx context.Context, orderID model.OrderID) error {
	foundOrder, err := o.repOrder.FindByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}

		return err
	}

	if foundOrder.Status != model.OrderStatusNew {
		return ErrPayedOrder
	}

	err = o.transactor.RunRepeatableRead(ctx, func(txCtx context.Context) error {
		for _, foundOrderItem := range foundOrder.Items {
			if err = o.repStock.ShipReserve(txCtx, foundOrderItem.SKU, foundOrderItem.Count); err != nil {
				return ErrShipReserve
			}
		}

		return o.repOrder.ChangeStatus(txCtx, orderID, model.OrderStatusPayed)
	})

	if err == nil {
		if sendErr := o.statusSender.Send(foundOrder.ID, model.OrderStatusPayed); sendErr != nil {
			fmt.Println("Failed send payment order status: ", sendErr.Error())
		}
	}

	return err
}

func (o *Order) Cancel(ctx context.Context, orderID model.OrderID) error {
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

	err = o.transactor.RunRepeatableRead(ctx, func(txCtx context.Context) error {
		for _, foundOrderItem := range foundOrder.Items {
			if err = o.repStock.CancelReserve(txCtx, foundOrderItem.SKU, foundOrderItem.Count); err != nil {
				return ErrCancelReserve
			}
		}

		return o.repOrder.ChangeStatus(txCtx, orderID, model.OrderStatusCanceled)
	})

	if err == nil {
		if sendErr := o.statusSender.Send(foundOrder.ID, model.OrderStatusCanceled); sendErr != nil {
			fmt.Println("Failed send cancel order status: ", sendErr.Error())
		}
	}

	return err
}

func (o *Order) CancelUnpaidWithDuration(ctx context.Context, duration time.Duration) error {
	orders, err := o.repOrder.FindByUnpaidStatusWithDuration(ctx, duration)
	if err != nil {
		return err
	}

	for _, orderItem := range orders {
		if err = o.Cancel(ctx, orderItem.ID); err != nil {
			return err
		}
	}

	return nil
}
