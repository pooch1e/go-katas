/*
Heap katas — problems where "give me the smallest (or largest) so far" is
asked over and over, solved with the standard library's container/heap.
*/
package heapkata

import (
	"container/heap"
	"errors"
)

var ErrInvalidK = errors.New("k out of range")

// IntHeap is a min-heap of ints for container/heap. container/heap does the
// sifting; it only needs these five methods to reach your slice.
type IntHeap []int

var _ heap.Interface = (*IntHeap)(nil)

func (h IntHeap) Len() int {
	return 0
}

func (h IntHeap) Less(i, j int) bool {
	return false
}

func (h IntHeap) Swap(i, j int) {
}

func (h *IntHeap) Push(x any) {
}

func (h *IntHeap) Pop() any {
	return nil
}

func KthLargest(nums []int, k int) (int, error) {
	return 0, nil
}

func TopKFrequent(nums []int, k int) []int {
	return nil
}

func MergeKSorted(lists [][]int) []int {
	return nil
}
