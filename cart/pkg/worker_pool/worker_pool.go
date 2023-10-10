package worker_pool

import (
	"context"
	"fmt"
	"sync"
)

const defaultWorkerCount = 3

type workerPool[Task any, Result any] struct {
	workerCount int
	processing  func(Task) Result
}

func New[Task any, Result any](processing func(Task) Result) *workerPool[Task, Result] {
	return &workerPool[Task, Result]{
		workerCount: defaultWorkerCount,
		processing:  processing,
	}
}

func (wp *workerPool[Task, Result]) Run(ctx context.Context, input <-chan Task) <-chan Result {
	resultCh := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(wp.workerCount)

	for i := 0; i < wp.workerCount; i++ {
		go func(index int) {
			fmt.Printf("Runned worker #%d\n", index)
			defer func() {
				fmt.Printf("Finished worker #%d\n", index)
				wg.Done()
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-input:
					if !ok {
						return
					}

					fmt.Printf("Worker #%d received %v\n", index, data)

					select {
					case <-ctx.Done():
						return
					case resultCh <- wp.processing(data):
						fmt.Printf("Worker #%d return %v\n", index, wp.processing(data))
					}
				}
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func (wp *workerPool[Task, Result]) SetWorkerCount(count int) {
	wp.workerCount = count
}
