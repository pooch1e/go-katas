/*
CompressRuns tests — run-length encode a string: "aaabcc" becomes "a3bc2".
A run of one is written as just the character.
*/
package stringkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	stringkata "github.com/pooch1e/go-katas/language/04-strings"
)

func TestCompressRuns(t *testing.T) {
	t.Run("compresses repeated characters", func(t *testing.T) {
		assert.Equal(t, "a3bc2", stringkata.CompressRuns("aaabcc"))
	})

	t.Run("leaves a string with no runs unchanged", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("returns an empty string for empty input", func(t *testing.T) {
		t.Skip("todo")
	})

	// "a12" should mean twelve a's, not a1 followed by a2.
	t.Run("handles runs of ten or more", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("treats separate runs of the same character separately", func(t *testing.T) {
		t.Skip("todo")
	})
}
