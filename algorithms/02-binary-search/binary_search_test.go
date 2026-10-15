/*
Binary search tests — find, find-where-it-would-go, find-the-boundary, and
search a sorted slice that's been rotated.
*/
package searchkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	searchkata "github.com/pooch1e/go-katas/algorithms/02-binary-search"
)

func TestBinarySearch(t *testing.T) {
	xs := []int{1, 3, 5, 7, 9, 11}

	t.Run("finds every element", func(t *testing.T) {
		for i, v := range xs {
			assert.Equal(t, i, searchkata.BinarySearch(xs, v))
		}
	})

	t.Run("returns -1 for a missing value", func(t *testing.T) {
		assert.Equal(t, -1, searchkata.BinarySearch(xs, 4))
	})

	t.Run("returns -1 below the first and above the last element", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns -1 for an empty slice", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestSearchInsert(t *testing.T) {
	t.Run("returns the index of an existing value", func(t *testing.T) {
		assert.Equal(t, 2, searchkata.SearchInsert([]int{1, 3, 5, 6}, 5))
	})

	t.Run("returns where a missing value would be inserted", func(t *testing.T) {
		assert.Equal(t, 1, searchkata.SearchInsert([]int{1, 3, 5, 6}, 2))
	})

	t.Run("returns len for a value bigger than everything", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns 0 for a value smaller than everything", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestFirstBadVersion(t *testing.T) {
	// Each call to isBad is "expensive". A linear scan would call it up to
	// n times; binary search needs about log2(n).
	t.Run("finds the first bad version with few checks", func(t *testing.T) {
		calls := 0
		isBad := func(v int) bool { calls++; return v >= 1_234_567 }

		assert.Equal(t, 1_234_567, searchkata.FirstBadVersion(2_000_000, isBad))
		assert.LessOrEqual(t, calls, 25)
	})

	t.Run("version 1 is bad", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("only the last version is bad", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestSearchRotated(t *testing.T) {
	// [4 5 6 7 0 1 2] is [0 1 2 4 5 6 7] rotated. At least one half of any
	// split is still sorted. Work out which, then ask: is the target in it?
	t.Run("finds a value in the rotated part", func(t *testing.T) {
		assert.Equal(t, 5, searchkata.SearchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 1))
	})

	t.Run("finds a value before the rotation point", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns -1 for a missing value", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("works on an unrotated slice", func(t *testing.T) {
		t.Skip("todo")
	})
}
