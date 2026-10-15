/*
Set kata tests — single-pass solutions using a seen-set.
*/
package setkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	setkata "github.com/pooch1e/go-katas/language/07-sets"
)

func TestHasDuplicates(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want bool
	}{
		{"no duplicates", []int{1, 2, 3}, false},
		{"one duplicate", []int{1, 2, 1}, true},
		{"empty", []int{}, false},
		{"duplicate at the very end", []int{5, 4, 3, 2, 1, 5}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, setkata.HasDuplicates(tc.in))
		})
	}

	t.Run("works for strings too", func(t *testing.T) {
		assert.True(t, setkata.HasDuplicates([]string{"a", "b", "a"}))
	})
}

func TestCountUniqueWords(t *testing.T) {
	t.Run("counts distinct words ignoring case and punctuation", func(t *testing.T) {
		assert.Equal(t, 3, setkata.CountUniqueWords("Go, go GO! Gopher go-er"))
	})

	t.Run("returns zero for blank text", func(t *testing.T) {
		assert.Equal(t, 0, setkata.CountUniqueWords("   "))
	})
}

func TestFirstRepeatedChar(t *testing.T) {
	t.Run("finds the first character seen twice", func(t *testing.T) {
		got, ok := setkata.FirstRepeatedChar("abcdbea")
		assert.True(t, ok)
		assert.Equal(t, 'b', got)
	})

	t.Run("reports false when nothing repeats", func(t *testing.T) {
		_, ok := setkata.FirstRepeatedChar("abc")
		assert.False(t, ok)
	})
}

func TestCommonElements(t *testing.T) {
	t.Run("returns values in both slices, in a's order, without duplicates", func(t *testing.T) {
		assert.Equal(t, []int{3, 1}, setkata.CommonElements([]int{3, 1, 3, 7}, []int{1, 3, 9}))
	})

	t.Run("returns an empty slice when nothing is shared", func(t *testing.T) {
		t.Skip("todo")
	})
}

// The nested-loop answer is O(n²). With a set of numbers seen so far you
// can do it in one pass: for each n, is target-n already in the set?
func TestHasPairWithSum(t *testing.T) {
	t.Run("finds a pair that adds up to the target", func(t *testing.T) {
		assert.True(t, setkata.HasPairWithSum([]int{8, 3, 5, 1}, 9))
	})

	t.Run("does not use the same element twice", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("can use two equal elements", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns false when no pair works", func(t *testing.T) {
		t.Skip("todo")
	})
}
