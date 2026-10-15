/*
Graph tests — DFS, BFS on a grid, flood fill, topological sort and Dijkstra.
*/
package graphkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	graphkata "github.com/pooch1e/go-katas/algorithms/05-graphs"
)

func TestHasPath(t *testing.T) {
	g := graphkata.Graph{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"d"},
		"d": {},
		"e": {"a"},
	}

	t.Run("follows edges to a reachable node", func(t *testing.T) {
		assert.True(t, graphkata.HasPath(g, "a", "d"))
	})

	t.Run("edges are one-way", func(t *testing.T) {
		assert.False(t, graphkata.HasPath(g, "d", "a"))
	})

	// Without a visited set, a cycle sends DFS round forever.
	t.Run("terminates on a graph with a cycle", func(t *testing.T) {
		cyclic := graphkata.Graph{"x": {"y"}, "y": {"x"}, "z": {}}
		assert.False(t, graphkata.HasPath(cyclic, "x", "z"))
	})

	t.Run("a node reaches itself", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestShortestPathGrid(t *testing.T) {
	grid := []string{
		"..#....",
		"..#.##.",
		"....#..",
		"###.#.#",
		"......."}

	t.Run("counts the fewest steps around walls", func(t *testing.T) {
		steps, ok := graphkata.ShortestPathGrid(grid, graphkata.Cell{Row: 0, Col: 0}, graphkata.Cell{Row: 0, Col: 6})
		require.True(t, ok)
		assert.Equal(t, 10, steps)
	})

	t.Run("start equals end is zero steps", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("reports false when the end is walled off", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestNumIslands(t *testing.T) {
	t.Run("counts groups of connected land (no diagonals)", func(t *testing.T) {
		grid := []string{
			"11000",
			"11010",
			"00100",
			"00011"}
		assert.Equal(t, 4, graphkata.NumIslands(grid))
	})

	t.Run("all water is zero islands", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("one big island", func(t *testing.T) {
		t.Skip("todo")
	})
}

// A valid order has every dependency before the thing that needs it. Many
// orders can be valid, so check the rule rather than one exact answer.
func assertValidOrder(t *testing.T, dependsOn graphkata.Graph, order []string) {
	t.Helper()
	pos := map[string]int{}
	for i, n := range order {
		pos[n] = i
	}
	for n, deps := range dependsOn {
		require.Contains(t, pos, n)
		for _, d := range deps {
			require.Contains(t, pos, d)
			assert.Less(t, pos[d], pos[n], "%s must come before %s", d, n)
		}
	}
}

func TestTopologicalSort(t *testing.T) {
	t.Run("orders every dependency first", func(t *testing.T) {
		dependsOn := graphkata.Graph{
			"app":    {"http", "db"},
			"http":   {"log"},
			"db":     {"log", "config"},
			"log":    {"config"},
			"config": {},
		}
		order, err := graphkata.TopologicalSort(dependsOn)
		require.NoError(t, err)
		assert.Len(t, order, 5)
		assertValidOrder(t, dependsOn, order)
	})

	t.Run("returns ErrCycle for a circular dependency", func(t *testing.T) {
		_, err := graphkata.TopologicalSort(graphkata.Graph{"a": {"b"}, "b": {"c"}, "c": {"a"}})
		assert.ErrorIs(t, err, graphkata.ErrCycle)
	})

	t.Run("includes nodes that only appear as dependencies", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("an empty graph gives an empty order", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestDijkstra(t *testing.T) {
	g := graphkata.WeightedGraph{
		"A": {{To: "B", Weight: 4}, {To: "C", Weight: 1}},
		"C": {{To: "B", Weight: 2}, {To: "D", Weight: 7}},
		"B": {{To: "D", Weight: 1}},
		"D": {},
		"E": {},
	}

	// The direct A→B edge costs 4, but A→C→B costs 3. Dijkstra must not
	// lock in a distance before every cheaper route has had its chance.
	t.Run("finds the cheapest distance to every reachable node", func(t *testing.T) {
		want := map[string]int{"A": 0, "C": 1, "B": 3, "D": 4}
		assert.Equal(t, want, graphkata.Dijkstra(g, "A"))
	})

	t.Run("unreachable nodes are left out of the result", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("once heaps are done: rewrite using container/heap and keep this green", func(t *testing.T) {
		t.Skip("todo")
	})
}
