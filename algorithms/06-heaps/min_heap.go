/*
MinHeap[T] is a binary heap stored in a slice: the parent of index i is at
(i-1)/2 and its children at 2i+1 and 2i+2. The smallest value is always at
index 0.
*/
package heapkata

import "cmp"

type MinHeap[T cmp.Ordered] struct {
	items []T
}

func (h *MinHeap[T]) Push(v T) {
}

func (h *MinHeap[T]) Pop() (T, bool) {
	var zero T
	return zero, false
}

func (h *MinHeap[T]) Peek() (T, bool) {
	var zero T
	return zero, false
}

func (h *MinHeap[T]) Len() int {
	return 0
}
