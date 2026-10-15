/*
CaesarCipher tests — shift every letter along the alphabet, wrapping round,
keeping case, and leaving everything else alone.
*/
package stringkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	stringkata "github.com/pooch1e/go-katas/language/04-strings"
)

func TestCaesarCipher(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		shift int
		want  string
	}{
		{"shifts lowercase", "abc", 1, "bcd"},
		{"wraps past z", "xyz", 3, "abc"},
		{"keeps uppercase", "Hello", 1, "Ifmmp"},
		{"leaves non-letters alone", "a-b c!", 1, "b-c d!"},
		// In Go, -1 % 26 is -1, not 25. Same as JS, which surprises people.
		{"negative shift", "abc", -1, "zab"},
		{"shift larger than the alphabet", "abc", 27, "bcd"},
		{"shift of zero", "go", 0, "go"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, stringkata.CaesarCipher(tc.in, tc.shift))
		})
	}
}
