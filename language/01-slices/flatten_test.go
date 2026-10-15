/*
Flatten tests — turn a [][]int into a single []int, in order.
*/
package slicekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	slicekata "github.com/pooch1e/go-katas/language/01-slices"
)

func TestFlatten(t *testing.T) {
	t.Run("joins inner slices in order", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, slicekata.Flatten([][]int{{1, 2}, {3}, {4, 5}}))
	})

	t.Run("skips empty inner slices", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, slicekata.Flatten([][]int{{}, {1}, {}, {2}}))
	})

	t.Run("returns an empty slice when there is nothing to flatten", func(t *testing.T) {
		assert.Empty(t, slicekata.Flatten([][]int{}))
	})
}
