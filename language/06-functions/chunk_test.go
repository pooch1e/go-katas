/*
Chunk tests — split a slice into groups of at most size elements.
*/
package funckata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	funckata "github.com/pooch1e/go-katas/language/06-functions"
)

func TestChunk(t *testing.T) {
	t.Run("splits evenly", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2}, {3, 4}}, funckata.Chunk([]int{1, 2, 3, 4}, 2))
	})

	t.Run("puts the remainder in a shorter last chunk", func(t *testing.T) {
		assert.Equal(t, [][]int{{1, 2, 3}, {4, 5}}, funckata.Chunk([]int{1, 2, 3, 4, 5}, 3))
	})

	t.Run("returns no chunks for empty input", func(t *testing.T) {
		assert.Empty(t, funckata.Chunk([]int{}, 3))
	})

	// If your chunks are sub-slices of xs (xs[i:j]), appending to one can
	// overwrite the next. Module 01's RemoveAt trap again. Full slice
	// expressions, xs[i:j:j], cap the capacity to stop that.
	t.Run("appending to a chunk does not overwrite the next chunk", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("a size of zero or less returns nil", func(t *testing.T) {
		t.Skip("todo")
	})
}
