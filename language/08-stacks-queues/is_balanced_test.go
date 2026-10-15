/*
IsBalanced tests — is every bracket closed by the right partner, in the
right order? Non-bracket characters are ignored.
*/
package stackkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	stackkata "github.com/pooch1e/go-katas/language/08-stacks-queues"
)

func TestIsBalanced(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"()", true},
		{"([]{})", true},
		{"func() { x[0] }", true},
		{"(", false},
		{")", false},
		{"(]", false},
		// Counting alone says this is fine: one of each. Order matters.
		{"([)]", false},
		{"(()", false},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.want, stackkata.IsBalanced(tc.in))
		})
	}
}
