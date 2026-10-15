/*
Sorting tests — the same cases run against all four sorts. Bubble,
Insertion and Quick sort in place; MergeSort returns a new slice.
*/
package sortkata_test

import (
	"math/rand/v2"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	sortkata "github.com/pooch1e/go-katas/algorithms/01-sorting"
)

var cases = []struct {
	name string
	in   []int
}{
	{"empty", []int{}},
	{"single", []int{1}},
	{"already sorted", []int{1, 2, 3, 4, 5}},
	{"reversed", []int{5, 4, 3, 2, 1}},
	{"duplicates", []int{3, 1, 3, 2, 1}},
	{"negatives", []int{0, -5, 3, -1}},
}

var inPlace = map[string]func([]int){
	"BubbleSort":    sortkata.BubbleSort[int],
	"InsertionSort": sortkata.InsertionSort[int],
	"QuickSort":     sortkata.QuickSort[int],
	"MergeSort": func(xs []int) {
		copy(xs, sortkata.MergeSort(xs))
	},
}

func TestSorts(t *testing.T) {
	for sortName, sort := range inPlace {
		for _, tc := range cases {
			t.Run(sortName+"/"+tc.name, func(t *testing.T) {
				got := slices.Clone(tc.in)
				want := slices.Clone(tc.in)
				slices.Sort(want)

				sort(got)
				assert.Equal(t, want, got)
			})
		}
	}
}

func TestMergeSort(t *testing.T) {
	t.Run("returns a new slice and leaves the input alone", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("sorts strings", func(t *testing.T) {
		t.Skip("todo")
	})
}

// A big random input catches off-by-one errors that six tiny cases miss.
// Compare each sort against slices.Sort on 10,000 random ints.
func TestSortsOnRandomInput(t *testing.T) {
	t.Skip("todo")
}

// Benchmarks: go test -bench=. ./algorithms/01-sorting/
// Watch Bubble and Insertion fall off a cliff as n grows.
func BenchmarkSorts(b *testing.B) {
	for _, n := range []int{100, 1_000, 10_000} {
		input := rand.Perm(n)
		for name, sort := range inPlace {
			b.Run(name+"/"+strconv.Itoa(n), func(b *testing.B) {
				for b.Loop() {
					sort(slices.Clone(input))
				}
			})
		}
	}
}
