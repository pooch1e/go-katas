/*
FindMax tests — find the largest value, and return an error rather than a
made-up value when there isn't one.
*/
package slicekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	slicekata "github.com/pooch1e/go-katas/language/01-slices"
)

func TestFindMax(t *testing.T) {
	t.Run("returns the only element of a single element slice", func(t *testing.T) {
		got, err := slicekata.FindMax([]int{7})
		require.NoError(t, err)
		assert.Equal(t, 7, got)
	})

	t.Run("returns the largest element", func(t *testing.T) {
		got, err := slicekata.FindMax([]int{3, 42, 17, 8})
		require.NoError(t, err)
		assert.Equal(t, 42, got)
	})

	// If you start your "best so far" at 0, this one fails. What's a safer
	// starting value?
	t.Run("works when every element is negative", func(t *testing.T) {
		got, err := slicekata.FindMax([]int{-9, -2, -15})
		require.NoError(t, err)
		assert.Equal(t, -2, got)
	})

	// JS would hand back undefined or -Infinity. Go has neither, so the
	// function returns a second value saying "this didn't work".
	t.Run("returns ErrEmpty for an empty slice", func(t *testing.T) {
		_, err := slicekata.FindMax([]int{})
		assert.ErrorIs(t, err, slicekata.ErrEmpty)
	})
}
