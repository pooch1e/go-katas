/*
Reverse tests — reverse a string by characters, not bytes.
*/
package stringkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	stringkata "github.com/pooch1e/go-katas/language/04-strings"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name, in, want string
	}{
		{"plain ascii", "golang", "gnalog"},
		{"empty string", "", ""},
		{"single character", "x", "x"},
		{"keeps spaces", "ab cd", "dc ba"},
		// "é" is two bytes in UTF-8. Reversing bytes splits it into garbage.
		{"accented characters survive", "café", "éfac"},
		{"emoji survive", "go🚀", "🚀og"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, stringkata.Reverse(tc.in))
		})
	}
}
