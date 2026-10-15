/*
SumParallel tests — split the slice into chunks, sum each chunk in its own
goroutine, and combine the results.
*/
package concurrencykata_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	concurrencykata "github.com/pooch1e/go-katas/language/12-concurrency"
)

func TestSumParallel(t *testing.T) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i + 1
	}

	for _, workers := range []int{1, 2, 3, 7, 1000} {
		t.Run(fmt.Sprintf("%d workers match a plain loop", workers), func(t *testing.T) {
			assert.Equal(t, 500500, concurrencykata.SumParallel(nums, workers))
		})
	}

	t.Run("an empty slice sums to zero", func(t *testing.T) {
		t.Skip("todo")
	})

	// Run this kata with `go test -race ./language/12-concurrency/`. If two
	// goroutines write the same variable unprotected, the race detector
	// fails the test even when the answer happens to come out right.
	t.Run("more workers than numbers still works", func(t *testing.T) {
		t.Skip("todo")
	})
}
