package order

import (
	"context"
	"log"
	"time"
)

type Canceler interface {
	CancelUnpaidWithDuration(ctx context.Context, duration time.Duration) error
}

type CancelUnpaidWorker struct {
	canceler Canceler
	duration time.Duration
}

func NewCancelUnpaidWorker(canceler Canceler, duration time.Duration) *CancelUnpaidWorker {
	return &CancelUnpaidWorker{canceler: canceler, duration: duration}
}

func (w *CancelUnpaidWorker) Run(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Minute)
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
