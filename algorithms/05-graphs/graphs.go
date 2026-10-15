/*
Package graphkata covers graphs as adjacency maps and grids: depth-first and
breadth-first search, connected components, topological ordering, and
shortest weighted paths.
*/
package graphkata

import "errors"

var ErrCycle = errors.New("graph has a cycle")

type Graph map[string][]string

func HasPath(g Graph, from, to string) bool {
	return false
}

type Cell struct {
	Row, Col int
}

func ShortestPathGrid(grid []string, start, end Cell) (int, bool) {
	return 0, false
}

func NumIslands(grid []string) int {
	return 0
}

func TopologicalSort(dependsOn Graph) ([]string, error) {
	return nil, nil
}

type Edge struct {
	To     string
	Weight int
}

type WeightedGraph map[string][]Edge

func Dijkstra(g WeightedGraph, source string) map[string]int {
	return nil
}
