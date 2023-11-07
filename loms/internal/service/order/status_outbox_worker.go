package order

import (
	"context"
	"go.uber.org/zap"
	"route256/loms/internal/repository"
	"time"
)

type StatusOutbox struct {
	rep          repository.OrderStatusOutbox
	statusSender StatusSender
	logger       *zap.Logger
}

func NewStatusOutbox(rep repository.OrderStatusOutbox, statusSender StatusSender, logger *zap.Logger) *StatusOutbox {
	return &StatusOutbox{
		rep:          rep,
		statusSender: statusSender,
		logger:       logger,
	}
}

func (w *StatusOutbox) Run(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.processing(ctx)
		case <-ctx.Done():
			w.logger.Info("Order status outbox worker stopped.")
			return nil
		}
	}
}

func (w *StatusOutbox) processing(ctx context.Context) {
	list, err := w.rep.FindUndelivered(ctx)
	if err != nil {
		w.logger.Error("Failed find order statuses from outbox.", zap.Error(err))
	}

	for _, item := range list {
		if sendErr := w.statusSender.Send(item.OrderID, item.OrderStatus); sendErr != nil {
			w.logger.Error("Failed send order status from outbox.", zap.Error(sendErr))
			continue
		}

		if markErr := w.rep.MarkAsDelivery(ctx, item.ID); markErr != nil {
			w.logger.Error("Failed mark as delivery order status from outbox.", zap.Error(markErr))
		}
	}
}
