/*
IsValidBST tests — does every node respect the ordering rule for the WHOLE
subtree, not just its direct children?
*/
package treekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	treekata "github.com/pooch1e/go-katas/language/11-trees"
)

func node(v int, left, right *treekata.TreeNode[int]) *treekata.TreeNode[int] {
	return &treekata.TreeNode[int]{Value: v, Left: left, Right: right}
}

func TestIsValidBST(t *testing.T) {
	t.Run("a valid tree", func(t *testing.T) {
		root := node(5, node(3, node(1, nil, nil), node(4, nil, nil)), node(8, nil, nil))
		assert.True(t, treekata.IsValidBST(root))
	})

	t.Run("a child on the wrong side", func(t *testing.T) {
		root := node(5, node(6, nil, nil), nil)
		assert.False(t, treekata.IsValidBST(root))
	})

	// 6 is bigger than its parent 3, so the local check passes. But it sits
	// in 5's LEFT subtree, where everything must be smaller than 5.
	t.Run("a grandchild breaking an ancestor's bound", func(t *testing.T) {
		root := node(5, node(3, nil, node(6, nil, nil)), node(8, nil, nil))
		assert.False(t, treekata.IsValidBST(root))
	})

	t.Run("an empty tree is valid", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("duplicate values are invalid", func(t *testing.T) {
		t.Skip("todo")
	})
}
