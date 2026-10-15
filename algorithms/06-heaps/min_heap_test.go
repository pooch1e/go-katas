/*
MinHeap tests — build a binary heap by hand: sift up on push, sift down on pop.
*/
package heapkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	heapkata "github.com/pooch1e/go-katas/algorithms/06-heaps"
)

func TestMinHeap(t *testing.T) {
	t.Run("pops values in ascending order", func(t *testing.T) {
		var h heapkata.MinHeap[int]
		for _, v := range []int{5, 3, 8, 1, 9, 2} {
			h.Push(v)
		}

		var got []int
		for h.Len() > 0 {
			v, _ := h.Pop()
			got = append(got, v)
		}
		assert.Equal(t, []int{1, 2, 3, 5, 8, 9}, got)
	})

	t.Run("Peek shows the minimum without removing it", func(t *testing.T) {
		var h heapkata.MinHeap[string]
		h.Push("m")
		h.Push("a")
		v, ok := h.Peek()
		assert.True(t, ok)
		assert.Equal(t, "a", v)
		assert.Equal(t, 2, h.Len())
	})

	t.Run("Pop and Peek on an empty heap report false", func(t *testing.T) {
		t.Skip("todo")
	})

	// Interleaving pushes and pops exercises sift-down on a heap that isn't
	// just "everything pushed, then everything popped".
	t.Run("stays correct with pushes and pops interleaved", func(t *testing.T) {
		t.Skip("todo: compare against a sorted slice over a few hundred random operations")
	})
}
