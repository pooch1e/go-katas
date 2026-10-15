/*
HasCycle tests — detect a loop in a chain of nodes using O(1) extra memory.
*/
package listkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	listkata "github.com/pooch1e/go-katas/language/10-linked-lists"
)

func chain(n int) []*listkata.Node[int] {
	nodes := make([]*listkata.Node[int], n)
	for i := range nodes {
		nodes[i] = &listkata.Node[int]{Value: i}
	}
	for i := 0; i < n-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	return nodes
}

func TestHasCycle(t *testing.T) {
	t.Run("a straight chain has no cycle", func(t *testing.T) {
		assert.False(t, listkata.HasCycle(chain(5)[0]))
	})

	t.Run("detects the tail pointing back into the middle", func(t *testing.T) {
		nodes := chain(5)
		nodes[4].Next = nodes[2]
		assert.True(t, listkata.HasCycle(nodes[0]))
	})

	t.Run("a nil head has no cycle", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("detects a node pointing at itself", func(t *testing.T) {
		t.Skip("todo")
	})
}
