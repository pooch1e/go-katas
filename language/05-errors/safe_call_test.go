/*
SafeCall tests — run a function and turn any panic into an ordinary error.
This is the one place recover() belongs: at a boundary, so one bad call
can't crash the whole program.
*/
package errorkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	errorkata "github.com/pooch1e/go-katas/language/05-errors"
)

func TestSafeCall(t *testing.T) {
	t.Run("returns nil when the function returns normally", func(t *testing.T) {
		ran := false
		err := errorkata.SafeCall(func() { ran = true })
		require.NoError(t, err)
		assert.True(t, ran)
	})

	t.Run("converts a panic into ErrPanicked", func(t *testing.T) {
		err := errorkata.SafeCall(func() { panic("boom") })
		assert.ErrorIs(t, err, errorkata.ErrPanicked)
	})

	t.Run("includes the panic value in the message", func(t *testing.T) {
		t.Skip("todo")
	})

	// Runtime panics (nil map write, index out of range, divide by zero)
	// are recoverable the same way.
	t.Run("catches a runtime panic such as an index out of range", func(t *testing.T) {
		t.Skip("todo")
	})
}
