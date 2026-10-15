/*
MakeCounter tests — a closure that remembers state between calls.
*/
package funckata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	funckata "github.com/pooch1e/go-katas/language/06-functions"
)

func TestMakeCounter(t *testing.T) {
	t.Run("counts up from one", func(t *testing.T) {
		next := funckata.MakeCounter()
		assert.Equal(t, 1, next())
		assert.Equal(t, 2, next())
		assert.Equal(t, 3, next())
	})

	t.Run("each counter has its own state", func(t *testing.T) {
		t.Skip("todo")
	})
}
