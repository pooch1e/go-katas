/*
CountLetters tests — tally how often each letter appears, ignoring case and
anything that isn't a letter.
*/
package mapkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mapkata "github.com/pooch1e/go-katas/language/02-maps"
)

func TestCountLetters(t *testing.T) {
	t.Run("counts each letter", func(t *testing.T) {
		assert.Equal(t, map[rune]int{'b': 1, 'a': 3, 'n': 2}, mapkata.CountLetters("banana"))
	})

	t.Run("treats upper and lower case as the same letter", func(t *testing.T) {
		assert.Equal(t, map[rune]int{'g': 2, 'o': 2}, mapkata.CountLetters("GoGo"))
	})

	t.Run("ignores spaces, digits and punctuation", func(t *testing.T) {
		assert.Equal(t, map[rune]int{'h': 1, 'i': 1}, mapkata.CountLetters("h i 1!"))
	})

	// An empty map and a nil map print the same and both have len 0, but
	// assert.Equal tells them apart. Return a made map, not nil.
	t.Run("returns an empty map for an empty string", func(t *testing.T) {
		assert.Equal(t, map[rune]int{}, mapkata.CountLetters(""))
	})
}
