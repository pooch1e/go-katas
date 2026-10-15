/*
Package concurrencykata covers goroutines, channels, sync.WaitGroup,
select, and context for cancellation and timeouts — Go's answer to
Promise.all, Promise.race and AbortController.
*/
package concurrencykata

import (
	"context"
	"time"
)

func SumParallel(nums []int, workers int) int {
	return 0
}

func FetchAll(urls []string, fetch func(string) (string, error)) ([]string, error) {
	return nil, nil
}

func FirstResult(fns ...func() string) string {
	return ""
}

func WithTimeout[T any](d time.Duration, fn func(ctx context.Context) (T, error)) (T, error) {
	var zero T
	return zero, nil
}

func WorkerPool(jobs []int, workers int, work func(int) int) []int {
	return nil
}

func Retry(ctx context.Context, attempts int, delay time.Duration, fn func() error) error {
	return nil
}
