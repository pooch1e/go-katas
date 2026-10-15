/*
Invert tests — swap keys and values. Mostly yours to write.
*/
package mapkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mapkata "github.com/pooch1e/go-katas/language/02-maps"
)

func TestInvert(t *testing.T) {
	t.Run("swaps every key with its value", func(t *testing.T) {
		got, err := mapkata.Invert(map[string]string{"a": "x", "b": "y"})
		require.NoError(t, err)
		assert.Equal(t, map[string]string{"x": "a", "y": "b"}, got)
	})

	// Maps are reference types, like JS objects. Writing into m would change
	// the caller's map.
	t.Run("returns a new map and leaves the source alone", func(t *testing.T) {
		t.Skip("todo")
	})

	// Two keys with the same value means two entries fighting over one new
	// key. A plain m[k] = v silently overwrites. Check before you write.
	t.Run("returns ErrDuplicateValue when two keys share a value", func(t *testing.T) {
		t.Skip("todo")
	})
}
