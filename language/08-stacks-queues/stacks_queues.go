/*
Package stackkata uses plain slices as stacks (append / pop the end) and
queues (append / take the front), plus switch statements and strconv.
*/
package stackkata

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivideByZero      = errors.New("divide by zero")
	ErrNoPlayers         = errors.New("no players")
)

func IsBalanced(s string) bool {
	return false
}

func EvaluateRPN(tokens []string) (int, error) {
	return 0, nil
}

func HotPotato(names []string, passes int) (string, error) {
	return "", nil
}
