// Package worker_pool makes it easier to work on a more abstract worker pool.
package worker_pool

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/pkg/worker_pool"
)

type Processor func(ctx context.Context, task Task) *Result

type Task struct {
	SKU model.SKU
}

type Result struct {
	SKU         model.SKU
	ProductInfo model.ProductInfo
}

// WorkerPool is a wrapper, and provides a simpler interface.
type WorkerPool struct {
	workerPool *worker_pool.WorkerPool[Task, *Result]
	Tasks      chan Task
	Results    <-chan *Result
}

// New constructor.
func New() *WorkerPool {
	return &WorkerPool{}
}

// Run method initializes and starts a workers pool.
// Input tasks channel, return results channel.
func (wp *WorkerPool) Run(ctx context.Context, tasks <-chan Task, proc Processor) <-chan *Result {
	wp.workerPool = worker_pool.New(proc)
	return wp.workerPool.Run(ctx, tasks)
}
