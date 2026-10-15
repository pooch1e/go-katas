/*
CountOccurrences tests — count how many times a word appears in a slice.
*/
package slicekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	slicekata "github.com/pooch1e/go-katas/language/01-slices"
)

func TestCountOccurrences(t *testing.T) {
	words := []string{"go", "js", "go", "rust", "go"}

	t.Run("counts every match", func(t *testing.T) {
		assert.Equal(t, 3, slicekata.CountOccurrences(words, "go"))
	})

	t.Run("returns zero when there is no match", func(t *testing.T) {
		assert.Equal(t, 0, slicekata.CountOccurrences(words, "zig"))
	})

	t.Run("is case sensitive", func(t *testing.T) {
		assert.Equal(t, 0, slicekata.CountOccurrences(words, "Go"))
	})

	t.Run("returns zero for an empty slice", func(t *testing.T) {
		assert.Equal(t, 0, slicekata.CountOccurrences(nil, "go"))
	})
}
