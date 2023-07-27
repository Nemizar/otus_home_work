package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded    = errors.New("errors limit exceeded")
	ErrErrorsGoroutinesNumber = errors.New("errors goroutines number")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	if n <= 0 {
		return ErrErrorsGoroutinesNumber
	}

	ch := make(chan Task)

	var errCounter int32

	go func() {
		for _, t := range tasks {
			if atomic.LoadInt32(&errCounter) >= int32(m) {
				break
			}
			ch <- t
		}

		close(ch)
	}()

	wg := sync.WaitGroup{}

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				if task == nil {
					atomic.AddInt32(&errCounter, 1)

					continue
				}

				if err := task(); err != nil {
					atomic.AddInt32(&errCounter, 1)
				}
			}
		}()
	}

	wg.Wait()

	if errCounter > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
