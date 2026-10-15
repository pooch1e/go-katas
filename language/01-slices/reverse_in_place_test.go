/*
ReverseInPlace tests — reverse the caller's slice without allocating a new one.
*/
package slicekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	slicekata "github.com/pooch1e/go-katas/language/01-slices"
)

func TestReverseInPlace(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{"odd length", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"single element", []int{9}, []int{9}},
		{"empty", []int{}, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			slicekata.ReverseInPlace(tc.nums)
			assert.Equal(t, tc.want, tc.nums)
		})
	}
}
