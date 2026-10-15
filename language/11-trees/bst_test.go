/*
BST tests — insert, search, min/max, height, and two traversals.
*/
package treekata_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	treekata "github.com/pooch1e/go-katas/language/11-trees"
)

func sampleTree() *treekata.BST[int] {
	/*
	        8
	      /   \
	     3     10
	    / \      \
	   1   6      14
	      / \
	     4   7
	*/
	var tr treekata.BST[int]
	for _, v := range []int{8, 3, 10, 1, 6, 14, 4, 7} {
		tr.Insert(v)
	}
	return &tr
}

func TestBST(t *testing.T) {
	t.Run("Insert reports new values and rejects duplicates", func(t *testing.T) {
		var tr treekata.BST[int]
		assert.True(t, tr.Insert(5))
		assert.False(t, tr.Insert(5))
		assert.Equal(t, 1, tr.Len())
	})

	t.Run("Contains finds inserted values only", func(t *testing.T) {
		tr := sampleTree()
		assert.True(t, tr.Contains(7))
		assert.True(t, tr.Contains(14))
		assert.False(t, tr.Contains(5))
	})

	t.Run("Min and Max", func(t *testing.T) {
		tr := sampleTree()
		lo, ok := tr.Min()
		assert.True(t, ok)
		assert.Equal(t, 1, lo)
		hi, ok := tr.Max()
		assert.True(t, ok)
		assert.Equal(t, 14, hi)
	})

	t.Run("Min and Max of an empty tree report false", func(t *testing.T) {
		t.Skip("todo")
	})

	// Height counts nodes on the longest root-to-leaf path: 8-3-6-4.
	t.Run("Height of the sample tree is 4", func(t *testing.T) {
		assert.Equal(t, 4, sampleTree().Height())
	})

	t.Run("Height of an empty tree is 0 and of one node is 1", func(t *testing.T) {
		t.Skip("todo")
	})

	// Inserting already-sorted values builds a "tree" that's really a
	// linked list. What does that do to Height, and to Contains' speed?
	t.Run("Height of a tree built from sorted input", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("InOrder yields values in sorted order", func(t *testing.T) {
		assert.Equal(t, []int{1, 3, 4, 6, 7, 8, 10, 14}, slices.Collect(sampleTree().InOrder()))
	})

	t.Run("InOrder stops when the range loop breaks", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("LevelOrder groups values by depth, left to right", func(t *testing.T) {
		want := [][]int{{8}, {3, 10}, {1, 6, 14}, {4, 7}}
		assert.Equal(t, want, sampleTree().LevelOrder())
	})

	t.Run("works with strings", func(t *testing.T) {
		t.Skip("todo")
	})
}
