/*
GroupByLength tests — bucket words by their length, keeping input order
inside each bucket.
*/
package mapkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mapkata "github.com/pooch1e/go-katas/language/02-maps"
)

func TestGroupByLength(t *testing.T) {
	t.Run("groups words of the same length together", func(t *testing.T) {
		want := map[int][]string{2: {"go", "js"}, 4: {"rust"}}
		assert.Equal(t, want, mapkata.GroupByLength([]string{"go", "rust", "js"}))
	})

	t.Run("keeps the input order within each group", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns an empty map for no words", func(t *testing.T) {
		t.Skip("todo")
	})
}
