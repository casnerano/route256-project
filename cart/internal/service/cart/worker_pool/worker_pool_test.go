package worker_pool

import (
	"context"
	"github.com/stretchr/testify/assert"
	"route256/cart/internal/model"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestWorkerPool_Run(t *testing.T) {
	pool := New()
	tasksCh := make(chan Task)
	processing := func(ctx context.Context, task Task) *Result {
		return &Result{
			SKU: task.SKU,
			ProductInfo: model.ProductInfo{
				Name:  "Untitled",
				Price: 100,
			},
		}
	}

	results := pool.Run(context.Background(), tasksCh, processing)

	go func() {
		for result := range results {
			assert.Greater(t, result.SKU, model.SKU(0))
			assert.Greater(t, result.ProductInfo.Price, uint32(0))
			assert.NotEmpty(t, result.ProductInfo.Name)
		}
	}()

	tasksCh <- Task{SKU: 1}
	tasksCh <- Task{SKU: 2}
	tasksCh <- Task{SKU: 3}

	close(tasksCh)
}
