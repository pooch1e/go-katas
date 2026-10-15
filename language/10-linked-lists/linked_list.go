/*
LinkedList[T] is a singly linked list built from pointer-linked nodes. It
exposes its values as an iter.Seq so callers can range over it like a slice.
*/
package listkata

import "iter"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func FromSlice[T any](xs []T) *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) PushFront(v T) {
}

func (l *LinkedList[T]) PushBack(v T) {
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	var zero T
	return zero, false
}

func (l *LinkedList[T]) Len() int {
	return 0
}

func (l *LinkedList[T]) Reverse() {
}

func (l *LinkedList[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {}
}
