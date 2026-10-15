/*
Stack tests — a generic LIFO container. The zero value (Stack[int]{}) must
be ready to use with no constructor.
*/
package ifacekata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ifacekata "github.com/pooch1e/go-katas/language/09-interfaces-generics"
)

func TestStack(t *testing.T) {
	t.Run("pops in reverse push order", func(t *testing.T) {
		var s ifacekata.Stack[int]
		s.Push(1)
		s.Push(2)
		s.Push(3)

		for _, want := range []int{3, 2, 1} {
			got, ok := s.Pop()
			assert.True(t, ok)
			assert.Equal(t, want, got)
		}
	})

	t.Run("Pop on an empty stack reports false", func(t *testing.T) {
		var s ifacekata.Stack[string]
		_, ok := s.Pop()
		assert.False(t, ok)
	})

	t.Run("Peek returns the top without removing it", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("Len tracks pushes and pops", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("works with a struct type", func(t *testing.T) {
		t.Skip("todo")
	})
}
