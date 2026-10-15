/*
Retry tests — call fn until it succeeds, up to `attempts` times, waiting
`delay` between tries and giving up early if ctx is cancelled.
*/
package concurrencykata_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	concurrencykata "github.com/pooch1e/go-katas/language/12-concurrency"
)

func TestRetry(t *testing.T) {
	flaky := errors.New("flaky")

	t.Run("succeeds once fn stops failing", func(t *testing.T) {
		calls := 0
		err := concurrencykata.Retry(context.Background(), 5, time.Millisecond, func() error {
			calls++
			if calls < 3 {
				return flaky
			}
			return nil
		})
		assert.NoError(t, err)
		assert.Equal(t, 3, calls)
	})

	t.Run("returns the last error after running out of attempts", func(t *testing.T) {
		calls := 0
		err := concurrencykata.Retry(context.Background(), 3, time.Millisecond, func() error {
			calls++
			return flaky
		})
		assert.ErrorIs(t, err, flaky)
		assert.Equal(t, 3, calls)
	})

	// time.Sleep can't be interrupted. select on ctx.Done() and
	// time.After(delay) can.
	t.Run("stops waiting as soon as the context is cancelled", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("does not wait after the final attempt", func(t *testing.T) {
		t.Skip("todo")
	})
}
