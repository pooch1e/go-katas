/*
LinkedList tests — pushing, popping, reversing by relinking pointers, and
iterating with range-over-func.
*/
package listkata_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	listkata "github.com/pooch1e/go-katas/language/10-linked-lists"
)

func TestLinkedList(t *testing.T) {
	t.Run("the zero value is an empty list", func(t *testing.T) {
		var l listkata.LinkedList[int]
		assert.Equal(t, 0, l.Len())
		assert.Empty(t, slices.Collect(l.All()))
	})

	t.Run("PushBack appends in order", func(t *testing.T) {
		var l listkata.LinkedList[string]
		l.PushBack("a")
		l.PushBack("b")
		l.PushBack("c")
		assert.Equal(t, []string{"a", "b", "c"}, slices.Collect(l.All()))
		assert.Equal(t, 3, l.Len())
	})

	t.Run("PushFront prepends", func(t *testing.T) {
		var l listkata.LinkedList[int]
		l.PushFront(1)
		l.PushFront(2)
		assert.Equal(t, []int{2, 1}, slices.Collect(l.All()))
	})

	t.Run("FromSlice builds a list in slice order", func(t *testing.T) {
		l := listkata.FromSlice([]int{1, 2, 3})
		assert.Equal(t, []int{1, 2, 3}, slices.Collect(l.All()))
	})

	t.Run("PopFront removes and returns the head", func(t *testing.T) {
		l := listkata.FromSlice([]int{1, 2})
		v, ok := l.PopFront()
		assert.True(t, ok)
		assert.Equal(t, 1, v)
		assert.Equal(t, []int{2}, slices.Collect(l.All()))
	})

	t.Run("PopFront on an empty list reports false", func(t *testing.T) {
		var l listkata.LinkedList[int]
		_, ok := l.PopFront()
		assert.False(t, ok)
	})

	// Keeping a tail pointer makes PushBack O(1). Popping the last node has
	// to clear it too, or the next PushBack links onto a ghost.
	t.Run("PushBack still works after popping everything", func(t *testing.T) {
		l := listkata.FromSlice([]int{1})
		l.PopFront()
		l.PushBack(2)
		assert.Equal(t, []int{2}, slices.Collect(l.All()))
	})

	t.Run("Reverse relinks the nodes", func(t *testing.T) {
		l := listkata.FromSlice([]int{1, 2, 3, 4})
		l.Reverse()
		assert.Equal(t, []int{4, 3, 2, 1}, slices.Collect(l.All()))
	})

	t.Run("PushBack after Reverse appends to the new end", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("Reverse of an empty or single-node list is a no-op", func(t *testing.T) {
		t.Skip("todo")
	})

	// range-over-func: the loop body runs inside yield. break makes yield
	// return false, and your iterator must stop when that happens.
	t.Run("All supports breaking out of a range loop early", func(t *testing.T) {
		l := listkata.FromSlice([]int{1, 2, 3, 4})
		var seen []int
		for v := range l.All() {
			if v == 3 {
				break
			}
			seen = append(seen, v)
		}
		assert.Equal(t, []int{1, 2}, seen)
	})
}
