/*
WithTimeout tests — run fn with a context that is cancelled after d, and
return context.DeadlineExceeded if fn doesn't finish in time.
*/
package concurrencykata_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	concurrencykata "github.com/pooch1e/go-katas/language/12-concurrency"
)

func sleepThen(d time.Duration, v int) func(context.Context) (int, error) {
	return func(ctx context.Context) (int, error) {
		select {
		case <-time.After(d):
			return v, nil
		case <-ctx.Done():
			return 0, ctx.Err()
		}
	}
}

func TestWithTimeout(t *testing.T) {
	t.Run("returns the value when fn finishes in time", func(t *testing.T) {
		got, err := concurrencykata.WithTimeout(100*time.Millisecond, sleepThen(5*time.Millisecond, 42))
		require.NoError(t, err)
		assert.Equal(t, 42, got)
	})

	t.Run("returns DeadlineExceeded when fn is too slow", func(t *testing.T) {
		start := time.Now()
		_, err := concurrencykata.WithTimeout(20*time.Millisecond, sleepThen(time.Second, 42))
		assert.ErrorIs(t, err, context.DeadlineExceeded)
		assert.Less(t, time.Since(start), 200*time.Millisecond)
	})

	t.Run("passes through fn's own error", func(t *testing.T) {
		t.Skip("todo")
	})
}
