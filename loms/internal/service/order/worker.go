package order

import (
	"context"
	"log"
	"time"

	"route256/loms/internal/repository/memstore"
)

type Canceler interface {
	CancelUnpaidWithDuration(ctx context.Context, duration time.Duration) error
}

type cancelUnpaidOrderWorker struct {
	canceler Canceler
	duration time.Duration
}

// TODO: dependency injection with Canceler
func NewCancelUnpaidWorker(duration time.Duration) *cancelUnpaidOrderWorker {
	repStock := memstore.NewStockRepository()
	repOrder := memstore.NewOrderRepository()

	return &cancelUnpaidOrderWorker{canceler: New(repOrder, repStock), duration: duration}
}

func (w *cancelUnpaidOrderWorker) Run(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := w.canceler.CancelUnpaidWithDuration(ctx, w.duration); err != nil {
				log.Println("Failed cancel unpaid order:", err)
			}
		case <-ctx.Done():
			log.Println("Cancel unpaid order worker stopped.")
			return nil
		}
	}
}
