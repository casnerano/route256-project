package order

import (
	"context"
	"go.uber.org/zap"
	"time"
)

type Canceler interface {
	CancelUnpaidWithDuration(ctx context.Context, duration time.Duration) error
}

type CancelUnpaidWorker struct {
	canceler Canceler
	duration time.Duration
	logger   *zap.Logger
}

func NewCancelUnpaidWorker(canceler Canceler, duration time.Duration, logger *zap.Logger) *CancelUnpaidWorker {
	return &CancelUnpaidWorker{canceler: canceler, duration: duration, logger: logger}
}

func (w *CancelUnpaidWorker) Run(ctx context.Context) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := w.canceler.CancelUnpaidWithDuration(ctx, w.duration); err != nil {
				w.logger.Error("Failed cancel unpaid order.", zap.Error(err))
			}
		case <-ctx.Done():
			w.logger.Info("Cancel unpaid order worker stopped.")
			return nil
		}
	}
}
