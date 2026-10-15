/*
SumAll tests — add up every number in a slice.
*/
package slicekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	slicekata "github.com/pooch1e/go-katas/language/01-slices"
)

func TestSumAll(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{"empty slice sums to zero", []int{}, 0},
		{"nil slice sums to zero", nil, 0},
		{"single element", []int{7}, 7},
		{"several positives", []int{1, 2, 3, 4}, 10},
		{"mixed signs", []int{5, -3, 10, -2}, 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, slicekata.SumAll(tc.nums))
		})
	}
}
