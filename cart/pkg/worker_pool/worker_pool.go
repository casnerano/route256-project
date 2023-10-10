package worker_pool

import (
	"context"
	"fmt"
	"sync"
)

const defaultWorkerCount = 3

type WorkerPool[Task any, Result any] struct {
	workerCount int
	// callback function for processing incoming tasks
	processing func(context.Context, Task) Result
}

// New - worker pool constructor.
// The type parameter is used to set the Task type for the input channel
// and the Result type for the output channel.
func New[Task any, Result any](processing func(context.Context, Task) Result) *WorkerPool[Task, Result] {
	return &WorkerPool[Task, Result]{
		workerCount: defaultWorkerCount,
		processing:  processing,
	}
}

// The Run method starts all workers.
// Each worker is a goroutine that reads messages (task) from a single input channel,
// processes the message using a callback function and puts the result in the output channel.
func (wp *WorkerPool[Task, Result]) Run(ctx context.Context, tasks <-chan Task) <-chan Result {
	results := make(chan Result)

	// WaitGroup to control the completion of all workers to close the output channel.
	var wg sync.WaitGroup
	wg.Add(wp.workerCount)

	for i := 0; i < wp.workerCount; i++ {
		go func(index int) {
			fmt.Printf("Runned worker #%d\n", index)
			defer func() {
				fmt.Printf("Finished worker #%d\n", index)
				wg.Done()
			}()

			// Reads data (tasks) from the input channel
			// until the channel is closed or the context is canceled.

			for {
				select {
				// Context can be canceled.
				case <-ctx.Done():
					return

				// Reading data (tasks) from the input channel until the channel is closed.
				case task, ok := <-tasks:
					if !ok {
						return
					}

					fmt.Printf("Worker #%d received %v\n", index, task)

					// Processing task.
					result := wp.processing(ctx, task)

					// The result of processing the task is sent to the output channel.
					select {
					// When sending the result to the output channel, the context may be canceled,
					// since the consumer may not immediately read the message.
					case <-ctx.Done():
						return
					// Processing the message via callback and send the result in the output channel.
					case results <- result:
						fmt.Printf("Worker #%d sent result %v\n", index, result)
					}
				}
			}
		}(i + 1)
	}

	// WaitGroup waits for all goroutine to finish and closes the output channel.
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func (wp *WorkerPool[Task, Result]) SetWorkerCount(count int) {
	wp.workerCount = count
}
