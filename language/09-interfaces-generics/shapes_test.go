/*
Shape tests — implement two shapes, then write code that works with any Shape.
*/
package ifacekata_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	ifacekata "github.com/pooch1e/go-katas/language/09-interfaces-generics"
)

func TestShapes(t *testing.T) {
	t.Run("Rectangle area and perimeter", func(t *testing.T) {
		r := ifacekata.Rectangle{Width: 3, Height: 4}
		assert.InDelta(t, 12.0, r.Area(), 1e-9)
		assert.InDelta(t, 14.0, r.Perimeter(), 1e-9)
	})

	t.Run("Circle area and perimeter", func(t *testing.T) {
		c := ifacekata.Circle{Radius: 2}
		assert.InDelta(t, 4*math.Pi, c.Area(), 1e-9)
		assert.InDelta(t, 4*math.Pi, c.Perimeter(), 1e-9)
	})

	t.Run("TotalArea sums a mix of shapes", func(t *testing.T) {
		shapes := []ifacekata.Shape{
			ifacekata.Rectangle{Width: 2, Height: 5},
			ifacekata.Circle{Radius: 1},
		}
		assert.InDelta(t, 10+math.Pi, ifacekata.TotalArea(shapes), 1e-9)
	})

	t.Run("TotalArea of no shapes is zero", func(t *testing.T) {
		t.Skip("todo")
	})

	// Describe needs to know WHICH shape it got. That's what a type switch
	// is for: switch v := s.(type) { case Circle: ... }.
	t.Run("Describe names the concrete shape", func(t *testing.T) {
		assert.Equal(t, "circle r=2", ifacekata.Describe(ifacekata.Circle{Radius: 2}))
		assert.Equal(t, "rectangle 3x4", ifacekata.Describe(ifacekata.Rectangle{Width: 3, Height: 4}))
	})

	t.Run("Describe handles shapes it doesn't know about", func(t *testing.T) {
		t.Skip("todo: declare a tiny Shape type inside this test and pass it in")
	})
}
