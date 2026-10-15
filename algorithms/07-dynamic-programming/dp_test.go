/*
Dynamic programming tests. Every kata has a brute-force recursive answer
that's correct but exponential; the large cases only pass with a table.
*/
package dpkata_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	dpkata "github.com/pooch1e/go-katas/algorithms/07-dynamic-programming"
)

// You can climb 1 or 2 steps at a time. How many distinct ways to the top?
func TestClimbStairs(t *testing.T) {
	cases := []struct{ n, want int }{
		{1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 8},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, dpkata.ClimbStairs(tc.n), "ClimbStairs(%d)", tc.n)
	}

	t.Run("n = 80 finishes instantly", func(t *testing.T) {
		assert.Equal(t, 37889062373143906, dpkata.ClimbStairs(80))
	})

	// The table only ever looks back two rows. Can you drop it for two
	// variables and get O(1) space?
	t.Run("n = 0 has one way: do nothing", func(t *testing.T) {
		t.Skip("todo")
	})
}

// Fewest coins that add up to amount, or -1 if it can't be done.
func TestCoinChange(t *testing.T) {
	t.Run("uses the fewest coins", func(t *testing.T) {
		assert.Equal(t, 3, dpkata.CoinChange([]int{1, 2, 5}, 11))
	})

	// Greedy (always take the biggest coin that fits) gives 4+1+1 = 3
	// coins here. The real answer is 3+3 = 2.
	t.Run("beats the greedy answer", func(t *testing.T) {
		assert.Equal(t, 2, dpkata.CoinChange([]int{1, 3, 4}, 6))
	})

	t.Run("returns -1 when the amount can't be made", func(t *testing.T) {
		assert.Equal(t, -1, dpkata.CoinChange([]int{2}, 3))
	})

	t.Run("zero amount needs zero coins", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("a large amount finishes instantly", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"", "abc", 0},
	}
	for _, tc := range cases {
		t.Run(tc.a+"/"+tc.b, func(t *testing.T) {
			assert.Equal(t, tc.want, dpkata.LongestCommonSubsequence(tc.a, tc.b))
		})
	}

	t.Run("long inputs finish instantly", func(t *testing.T) {
		a := strings.Repeat("ab", 500)
		b := strings.Repeat("ba", 500)
		assert.Equal(t, 999, dpkata.LongestCommonSubsequence(a, b))
	})

	t.Run("is symmetric", func(t *testing.T) {
		t.Skip("todo")
	})
}

// 0/1 knapsack: each item is taken whole or not at all, at most once.
func TestKnapsack(t *testing.T) {
	t.Run("maximises value within the weight limit", func(t *testing.T) {
		weights := []int{1, 3, 4, 5}
		values := []int{1, 4, 5, 7}
		assert.Equal(t, 9, dpkata.Knapsack(weights, values, 7))
	})

	t.Run("nothing fits", func(t *testing.T) {
		t.Skip("todo")
	})

	// If your answer is too high, you're probably letting an item be
	// taken twice. Which direction you loop over capacity matters if you
	// use a single 1D row.
	t.Run("each item is used at most once", func(t *testing.T) {
		t.Skip("todo")
	})
}
