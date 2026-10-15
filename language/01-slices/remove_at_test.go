/*
RemoveAt tests — return a new slice without the element at index i. The
written tests show the happy path; the todo tests are yours to fill in.
*/
package slicekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	slicekata "github.com/pooch1e/go-katas/language/01-slices"
)

func TestRemoveAt(t *testing.T) {
	t.Run("removes the element at the given index", func(t *testing.T) {
		got, err := slicekata.RemoveAt([]int{10, 20, 30, 40}, 1)
		require.NoError(t, err)
		assert.Equal(t, []int{10, 30, 40}, got)
	})

	// The classic one-liner is append(nums[:i], nums[i+1:]...). It works,
	// but nums[:i] shares nums' backing array, so append writes straight
	// over the caller's data. Print the input after calling it and see.
	t.Run("does not mutate the input", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns ErrOutOfRange for a negative index", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns ErrOutOfRange for an index equal to len", func(t *testing.T) {
		t.Skip("todo")
	})
}
