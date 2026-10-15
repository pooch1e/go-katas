/*
FirstResult tests — Promise.race for Go: start everything, return whichever
finishes first.
*/
package concurrencykata_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	concurrencykata "github.com/pooch1e/go-katas/language/12-concurrency"
)

func after(d time.Duration, v string) func() string {
	return func() string {
		time.Sleep(d)
		return v
	}
}

func TestFirstResult(t *testing.T) {
	t.Run("returns the fastest function's value", func(t *testing.T) {
		got := concurrencykata.FirstResult(
			after(80*time.Millisecond, "slow"),
			after(10*time.Millisecond, "fast"),
			after(40*time.Millisecond, "medium"),
		)
		assert.Equal(t, "fast", got)
	})

	t.Run("does not wait for the slower functions", func(t *testing.T) {
		start := time.Now()
		concurrencykata.FirstResult(after(5*time.Millisecond, "a"), after(300*time.Millisecond, "b"))
		assert.Less(t, time.Since(start), 150*time.Millisecond)
	})

	// The losers still finish and try to send their result. On an
	// unbuffered channel nobody is listening, so they block forever: a
	// goroutine leak. How big should the channel's buffer be?
	t.Run("the losing goroutines can still finish", func(t *testing.T) {
		t.Skip("todo")
	})
}
