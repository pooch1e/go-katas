/*
DoubleAll tests — return a NEW slice with every number doubled, leaving the
input untouched.
*/
package slicekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	slicekata "github.com/pooch1e/go-katas/language/01-slices"
)

func TestDoubleAll(t *testing.T) {
	t.Run("doubles every element", func(t *testing.T) {
		assert.Equal(t, []int{2, 4, 6}, slicekata.DoubleAll([]int{1, 2, 3}))
	})

	t.Run("returns an empty slice for empty input", func(t *testing.T) {
		got := slicekata.DoubleAll([]int{})
		assert.NotNil(t, got, "return []int{}, not nil, so callers can tell 'no items' from 'no answer'")
		assert.Empty(t, got)
	})

	// A slice header points at a shared backing array. If you write into
	// nums directly, the caller's slice changes too, just like a JS array.
	t.Run("does not mutate the input", func(t *testing.T) {
		input := []int{1, 2, 3}
		slicekata.DoubleAll(input)
		assert.Equal(t, []int{1, 2, 3}, input)
	})
}
