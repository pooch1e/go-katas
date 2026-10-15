/*
Largest and Celsius tests — a constrained generic function, and a named
type that knows how to print itself.
*/
package ifacekata_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ifacekata "github.com/pooch1e/go-katas/language/09-interfaces-generics"
)

func TestLargest(t *testing.T) {
	t.Run("works on ints", func(t *testing.T) {
		got, ok := ifacekata.Largest([]int{3, 9, 2})
		assert.True(t, ok)
		assert.Equal(t, 9, got)
	})

	t.Run("works on strings, comparing alphabetically", func(t *testing.T) {
		got, ok := ifacekata.Largest([]string{"go", "zig", "c"})
		assert.True(t, ok)
		assert.Equal(t, "zig", got)
	})

	t.Run("reports false for an empty slice", func(t *testing.T) {
		_, ok := ifacekata.Largest([]float64{})
		assert.False(t, ok)
	})

	// Celsius's underlying type is float64, so it satisfies cmp.Ordered
	// too. Constraints like ~float64 match named types built on float64.
	t.Run("works on a named type", func(t *testing.T) {
		t.Skip("todo")
	})
}

func TestCelsius(t *testing.T) {
	// fmt.Sprint calls String() when a value implements fmt.Stringer.
	t.Run("prints with one decimal place and a unit", func(t *testing.T) {
		assert.Equal(t, "21.5°C", fmt.Sprint(ifacekata.Celsius(21.5)))
		assert.Equal(t, "-3.0°C", fmt.Sprint(ifacekata.Celsius(-3)))
	})
}
