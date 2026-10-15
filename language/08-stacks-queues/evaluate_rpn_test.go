/*
EvaluateRPN tests — evaluate Reverse Polish Notation: "3 4 +" means 3 + 4.
Division truncates toward zero like Go's integer /.
*/
package stackkata_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	stackkata "github.com/pooch1e/go-katas/language/08-stacks-queues"
)

func TestEvaluateRPN(t *testing.T) {
	valid := []struct {
		expr string
		want int
	}{
		{"42", 42},
		{"3 4 +", 7},
		{"10 4 -", 6},
		{"2 3 4 * +", 14},
		{"5 1 2 + 4 * + 3 -", 14},
		{"7 2 /", 3},
		{"-3 2 *", -6},
	}

	for _, tc := range valid {
		t.Run(tc.expr, func(t *testing.T) {
			got, err := stackkata.EvaluateRPN(strings.Fields(tc.expr))
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}

	t.Run("not enough operands is invalid", func(t *testing.T) {
		_, err := stackkata.EvaluateRPN([]string{"1", "+"})
		assert.ErrorIs(t, err, stackkata.ErrInvalidExpression)
	})

	t.Run("leftover operands is invalid", func(t *testing.T) {
		_, err := stackkata.EvaluateRPN([]string{"1", "2"})
		assert.ErrorIs(t, err, stackkata.ErrInvalidExpression)
	})

	t.Run("an unknown token is invalid", func(t *testing.T) {
		_, err := stackkata.EvaluateRPN([]string{"1", "2", "^"})
		assert.ErrorIs(t, err, stackkata.ErrInvalidExpression)
	})

	t.Run("dividing by zero returns ErrDivideByZero instead of panicking", func(t *testing.T) {
		_, err := stackkata.EvaluateRPN([]string{"1", "0", "/"})
		assert.ErrorIs(t, err, stackkata.ErrDivideByZero)
	})

	t.Run("an empty expression is invalid", func(t *testing.T) {
		_, err := stackkata.EvaluateRPN(nil)
		assert.ErrorIs(t, err, stackkata.ErrInvalidExpression)
	})
}
