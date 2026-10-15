/*
IsPalindrome tests — does it read the same both ways, ignoring case,
spaces and punctuation?
*/
package stringkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	stringkata "github.com/pooch1e/go-katas/language/04-strings"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"racecar", true},
		{"golang", false},
		{"", true},
		{"a", true},
		{"Racecar", true},
		{"A man, a plan, a canal: Panama", true},
		{"No 'x' in Nixon", true},
		{"ab", false},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.want, stringkata.IsPalindrome(tc.in))
		})
	}
}
