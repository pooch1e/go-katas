/*
Recursion and backtracking tests. Results whose order doesn't matter are
compared with assert.ElementsMatch.
*/
package recursionkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	recursionkata "github.com/pooch1e/go-katas/algorithms/04-recursion"
)

func TestFibonacci(t *testing.T) {
	cases := []struct{ n, want int }{
		{0, 0}, {1, 1}, {2, 1}, {10, 55}, {20, 6765},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, recursionkata.Fibonacci(tc.n), "Fibonacci(%d)", tc.n)
	}

	// Naive recursion makes about 2^90 calls here and never finishes. Run
	// with -timeout 10s so you find out quickly, then add a memo.
	t.Run("Fibonacci(90) finishes instantly", func(t *testing.T) {
		assert.Equal(t, 2880067194370816120, recursionkata.Fibonacci(90))
	})
}

func TestPermutations(t *testing.T) {
	t.Run("all orderings of three elements", func(t *testing.T) {
		want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
		assert.ElementsMatch(t, want, recursionkata.Permutations([]int{1, 2, 3}))
	})

	// If every result looks the same (or wrong), you appended your working
	// `path` slice itself. Every result then shares one backing array that
	// you keep overwriting. Copy it when you record a result.
	t.Run("results don't share memory", func(t *testing.T) {
		t.Skip("todo: mutate one returned permutation and check the others are untouched")
	})

	t.Run("one element has one permutation", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("n elements have n! permutations", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestSubsets(t *testing.T) {
	t.Run("every subset including empty and full", func(t *testing.T) {
		want := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}
		assert.ElementsMatch(t, want, recursionkata.Subsets([]int{1, 2, 3}))
	})

	t.Run("the empty set has one subset", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("n elements have 2^n subsets", func(t *testing.T) {
		t.Skip("todo")
	})
}

// Candidates may be reused any number of times. Each combination is listed
// in ascending order, and no combination appears twice.
func TestCombinationSum(t *testing.T) {
	t.Run("finds every combination that hits the target", func(t *testing.T) {
		want := [][]int{{2, 2, 3}, {7}}
		assert.ElementsMatch(t, want, recursionkata.CombinationSum([]int{2, 3, 6, 7}, 7))
	})

	// [2 3 3] and [3 2 3] are the same combination. Only ever choosing
	// candidates at or after the current index prevents the reorderings.
	t.Run("does not repeat combinations in a different order", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns no combinations when the target can't be reached", func(t *testing.T) {
		t.Skip("todo")
	})
}
