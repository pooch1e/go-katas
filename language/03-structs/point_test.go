/*
Point tests — value semantics: copying, == comparison, and methods that
return a new Point.
*/
package structkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	structkata "github.com/pooch1e/go-katas/language/03-structs"
)

func TestPoint(t *testing.T) {
	// Structs whose fields are all comparable can be compared with ==,
	// field by field. This passes already: nothing to implement.
	t.Run("two points with the same fields are equal", func(t *testing.T) {
		assert.True(t, structkata.Point{X: 1, Y: 2} == structkata.Point{X: 1, Y: 2})
	})

	t.Run("Add sums the coordinates", func(t *testing.T) {
		got := structkata.Point{X: 1, Y: 2}.Add(structkata.Point{X: 10, Y: 20})
		assert.Equal(t, structkata.Point{X: 11, Y: 22}, got)
	})

	t.Run("Moved returns a shifted copy", func(t *testing.T) {
		got := structkata.Point{X: 0, Y: 0}.Moved(3, -1)
		assert.Equal(t, structkata.Point{X: 3, Y: -1}, got)
	})

	// Assigning a struct copies it. Moving the copy must not touch the
	// original, which is the opposite of what a JS object would do.
	t.Run("Moved leaves the original alone", func(t *testing.T) {
		t.Skip("todo")
	})

	// Floats rarely compare exactly. Look up assert.InDelta.
	t.Run("DistanceTo uses Pythagoras", func(t *testing.T) {
		t.Skip("todo")
	})
}
