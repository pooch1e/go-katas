/*
Heap kata tests — container/heap in practice.
*/
package heapkata_test

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	heapkata "github.com/pooch1e/go-katas/algorithms/06-heaps"
)

func TestIntHeap(t *testing.T) {
	t.Run("works with container/heap", func(t *testing.T) {
		h := &heapkata.IntHeap{5, 2, 8}
		heap.Init(h)
		heap.Push(h, 1)
		assert.Equal(t, 1, heap.Pop(h))
		assert.Equal(t, 2, heap.Pop(h))
	})
}

// Keep a min-heap of the k largest seen so far. Its top is the answer.
// That's O(n log k), better than sorting everything when k is small.
func TestKthLargest(t *testing.T) {
	t.Run("finds the kth largest value", func(t *testing.T) {
		got, err := heapkata.KthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
		require.NoError(t, err)
		assert.Equal(t, 5, got)
	})

	t.Run("counts duplicates separately", func(t *testing.T) {
		got, err := heapkata.KthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
		require.NoError(t, err)
		assert.Equal(t, 4, got)
	})

	t.Run("returns ErrInvalidK for k < 1 or k > len", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestTopKFrequent(t *testing.T) {
	t.Run("returns the k most frequent values, most frequent first", func(t *testing.T) {
		assert.Equal(t, []int{1, 2}, heapkata.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	})

	// Without a tie-break rule the answer isn't deterministic.
	t.Run("breaks frequency ties by the smaller value", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("k larger than the number of distinct values returns them all", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestMergeKSorted(t *testing.T) {
	t.Run("merges several sorted slices", func(t *testing.T) {
		got := heapkata.MergeKSorted([][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}})
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, got)
	})

	t.Run("handles empty and uneven lists", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("no lists gives an empty slice", func(t *testing.T) {
		t.Skip("todo")
	})
}
