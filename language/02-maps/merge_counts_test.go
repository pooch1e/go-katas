/*
MergeCounts tests — combine two tallies, adding counts for shared keys.
*/
package mapkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mapkata "github.com/pooch1e/go-katas/language/02-maps"
)

func TestMergeCounts(t *testing.T) {
	t.Run("adds counts for keys in both maps", func(t *testing.T) {
		got := mapkata.MergeCounts(map[string]int{"go": 2}, map[string]int{"go": 3})
		assert.Equal(t, map[string]int{"go": 5}, got)
	})

	t.Run("keeps keys that appear in only one map", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("does not mutate either input", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("treats a nil map as empty", func(t *testing.T) {
		t.Skip("todo")
	})
}
