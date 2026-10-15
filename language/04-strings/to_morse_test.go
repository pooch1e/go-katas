/*
ToMorse tests — letters separated by a space, words separated by " / ".
*/
package stringkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	stringkata "github.com/pooch1e/go-katas/language/04-strings"
)

func TestToMorse(t *testing.T) {
	t.Run("encodes a single word", func(t *testing.T) {
		got, err := stringkata.ToMorse("SOS")
		require.NoError(t, err)
		assert.Equal(t, "... --- ...", got)
	})

	t.Run("separates words with a slash", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("accepts lowercase input", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns ErrUnsupportedChar for punctuation", func(t *testing.T) {
		t.Skip("todo")
	})
}
