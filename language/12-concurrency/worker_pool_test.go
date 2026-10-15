/*
WorkerPool tests — process every job using at most `workers` goroutines,
returning results in job order.
*/
package concurrencykata_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	concurrencykata "github.com/pooch1e/go-katas/language/12-concurrency"
)

func TestWorkerPool(t *testing.T) {
	double := func(n int) int { return n * 2 }

	t.Run("returns results in job order", func(t *testing.T) {
		got := concurrencykata.WorkerPool([]int{1, 2, 3, 4, 5}, 2, double)
		assert.Equal(t, []int{2, 4, 6, 8, 10}, got)
	})

	t.Run("never runs more than `workers` jobs at once", func(t *testing.T) {
		var running, peak atomic.Int32
		work := func(n int) int {
			now := running.Add(1)
			for {
				old := peak.Load()
				if now <= old || peak.CompareAndSwap(old, now) {
					break
				}
			}
			time.Sleep(10 * time.Millisecond)
			running.Add(-1)
			return n
		}

		concurrencykata.WorkerPool(make([]int, 20), 3, work)
		assert.Equal(t, int32(3), peak.Load())
	})

	t.Run("no jobs returns an empty slice", func(t *testing.T) {
		t.Skip("todo")
	})
}
