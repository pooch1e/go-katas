/*
HotPotato tests — players stand in a circle passing a potato. After
`passes` passes, whoever holds it is out. Repeat until one player is left.
*/
package stackkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	stackkata "github.com/pooch1e/go-katas/language/08-stacks-queues"
)

func TestHotPotato(t *testing.T) {
	// Passing moves the front player to the back. With passes=1 and
	// [A B C]: pass -> [B C A], B is out -> [C A], pass -> [A C], A is out.
	t.Run("returns the last player standing", func(t *testing.T) {
		got, err := stackkata.HotPotato([]string{"A", "B", "C"}, 1)
		require.NoError(t, err)
		assert.Equal(t, "C", got)
	})

	t.Run("a single player wins immediately", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("zero passes eliminates players in order", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns ErrNoPlayers for an empty circle", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("does not mutate the input", func(t *testing.T) {
		t.Skip("todo")
	})
}
