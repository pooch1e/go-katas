/*
TitleCase tests — capitalise the first letter of every word, lowercase the
rest, and collapse runs of whitespace to single spaces.
*/
package stringkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	stringkata "github.com/pooch1e/go-katas/language/04-strings"
)

func TestTitleCase(t *testing.T) {
	cases := []struct {
		name, in, want string
	}{
		{"lowercase words", "the go programming language", "The Go Programming Language"},
		{"shouting", "HELLO WORLD", "Hello World"},
		{"extra whitespace", "  too   many\tspaces ", "Too Many Spaces"},
		{"single letter words", "a b c", "A B C"},
		{"empty", "", ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, stringkata.TitleCase(tc.in))
		})
	}
}
