/*
Two pointer and sliding window tests.
*/
package pointerkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pointerkata "github.com/pooch1e/go-katas/algorithms/03-two-pointers"
)

func TestTwoSumSorted(t *testing.T) {
	t.Run("returns the indexes of a pair that sums to target", func(t *testing.T) {
		nums := []int{1, 3, 4, 6, 9}
		i, j, ok := pointerkata.TwoSumSorted(nums, 10)
		require.True(t, ok)
		assert.Equal(t, 10, nums[i]+nums[j])
		assert.Less(t, i, j)
	})

	t.Run("reports false when no pair exists", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("does not pair an element with itself", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestRemoveDuplicatesSorted(t *testing.T) {
	// LeetCode-style in-place compaction: the first k elements hold the
	// unique values; whatever sits after index k doesn't matter.
	t.Run("compacts unique values to the front and returns their count", func(t *testing.T) {
		nums := []int{1, 1, 2, 3, 3, 3, 4}
		k := pointerkata.RemoveDuplicatesSorted(nums)
		assert.Equal(t, 4, k)
		assert.Equal(t, []int{1, 2, 3, 4}, nums[:k])
	})

	t.Run("returns 0 for an empty slice", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("all the same value", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestMaxSumSubarrayK(t *testing.T) {
	t.Run("finds the best window of size k", func(t *testing.T) {
		got, err := pointerkata.MaxSumSubarrayK([]int{2, 1, 5, 1, 3, 2}, 3)
		require.NoError(t, err)
		assert.Equal(t, 9, got)
	})

	// Starting "best" at 0 is the FindMax bug from language/01 again.
	t.Run("works when every number is negative", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns ErrWindowTooBig when k > len", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestLongestUniqueSubstring(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.want, pointerkata.LongestUniqueSubstring(tc.in))
		})
	}

	// The left edge must only ever move RIGHT. In "abba", when the second
	// a arrives, its last position is behind the window. Don't jump back.
	t.Run("abba", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("counts runes, not bytes", func(t *testing.T) {
		t.Skip("todo")
	})
}
