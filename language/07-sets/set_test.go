/*
Set tests — build the set type the other katas lean on.
*/
package setkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	setkata "github.com/pooch1e/go-katas/language/07-sets"
)

func TestSet(t *testing.T) {
	t.Run("a new set holds its initial items", func(t *testing.T) {
		s := setkata.NewSet("go", "js")
		assert.True(t, s.Has("go"))
		assert.True(t, s.Has("js"))
		assert.False(t, s.Has("rust"))
		assert.Equal(t, 2, s.Len())
	})

	t.Run("duplicates in the initial items are counted once", func(t *testing.T) {
		assert.Equal(t, 2, setkata.NewSet(1, 1, 2).Len())
	})

	// Returning whether the item was new saves callers a separate Has check.
	t.Run("Add reports whether the item was new", func(t *testing.T) {
		s := setkata.NewSet[int]()
		assert.True(t, s.Add(5))
		assert.False(t, s.Add(5))
		assert.Equal(t, 1, s.Len())
	})

	t.Run("Remove deletes an item", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("Remove of a missing item is a no-op", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("Intersect returns items in both sets", func(t *testing.T) {
		t.Skip("todo")
	})
}
