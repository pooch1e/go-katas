/*
Map, Filter and Reduce tests — write the three classics yourself as generic
functions. After this you'll know exactly what JS's array methods do.
*/
package funckata_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	funckata "github.com/pooch1e/go-katas/language/06-functions"
)

func TestMap(t *testing.T) {
	t.Run("applies the function to every element", func(t *testing.T) {
		got := funckata.Map([]int{1, 2, 3}, func(n int) int { return n * n })
		assert.Equal(t, []int{1, 4, 9}, got)
	})

	t.Run("can change the element type", func(t *testing.T) {
		got := funckata.Map([]int{1, 20}, strconv.Itoa)
		assert.Equal(t, []string{"1", "20"}, got)
	})

	t.Run("returns an empty slice for empty input", func(t *testing.T) {
		got := funckata.Map([]int{}, func(n int) int { return n })
		assert.NotNil(t, got)
		assert.Empty(t, got)
	})
}

func TestFilter(t *testing.T) {
	t.Run("keeps elements where the predicate is true", func(t *testing.T) {
		got := funckata.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
		assert.Equal(t, []int{2, 4}, got)
	})

	t.Run("returns an empty slice when nothing matches", func(t *testing.T) {
		got := funckata.Filter([]string{"a", "b"}, func(s string) bool { return false })
		assert.NotNil(t, got)
		assert.Empty(t, got)
	})
}

func TestReduce(t *testing.T) {
	t.Run("folds a slice into a single value", func(t *testing.T) {
		got := funckata.Reduce([]int{1, 2, 3, 4}, 0, func(acc, n int) int { return acc + n })
		assert.Equal(t, 10, got)
	})

	t.Run("the accumulator can be a different type", func(t *testing.T) {
		got := funckata.Reduce([]string{"go", "rust"}, 0, func(acc int, s string) int { return acc + len(s) })
		assert.Equal(t, 6, got)
	})

	t.Run("returns the initial value for empty input", func(t *testing.T) {
		got := funckata.Reduce([]int{}, 42, func(acc, n int) int { return acc + n })
		assert.Equal(t, 42, got)
	})
}
