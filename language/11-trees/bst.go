/*
BST[T] is a binary search tree: every value in a node's left subtree is
smaller, every value in its right subtree is larger. Recursion does most of
the work.
*/
package treekata

import (
	"cmp"
	"iter"
)

type TreeNode[T cmp.Ordered] struct {
	Value       T
	Left, Right *TreeNode[T]
}

type BST[T cmp.Ordered] struct {
	root *TreeNode[T]
	size int
}

func (t *BST[T]) Insert(v T) bool {
	return false
}

func (t *BST[T]) Contains(v T) bool {
	return false
}

func (t *BST[T]) Len() int {
	return 0
}

func (t *BST[T]) Min() (T, bool) {
	var zero T
	return zero, false
}

func (t *BST[T]) Max() (T, bool) {
	var zero T
	return zero, false
}

func (t *BST[T]) Height() int {
	return 0
}

func (t *BST[T]) InOrder() iter.Seq[T] {
	return func(yield func(T) bool) {}
}

func (t *BST[T]) LevelOrder() [][]T {
	return nil
}
