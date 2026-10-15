/*
Package mapkata covers maps: creating them with make or a literal, the
comma-ok lookup, iterating (in random order), and maps as reference types.
*/
package mapkata

import "errors"

var ErrDuplicateValue = errors.New("duplicate value")

func CountLetters(s string) map[rune]int {
	return nil
}

func GroupByLength(words []string) map[int][]string {
	return nil
}

func Invert(m map[string]string) (map[string]string, error) {
	return nil, nil
}

func LookupAge(ages map[string]int, name string, fallback int) int {
	return 0
}

func MergeCounts(a, b map[string]int) map[string]int {
	return nil
}
