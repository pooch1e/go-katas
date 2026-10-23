/*
Package mapkata covers maps: creating them with make or a literal, the
comma-ok lookup, iterating (in random order), and maps as reference types.
*/
package mapkata

import (
	"errors"
	"unicode"
)

var ErrDuplicateValue = errors.New("duplicate value")

func CountLetters(s string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		count[unicode.ToLower(c)]++
	}
	return count
}

func GroupByLength(words []string) map[int][]string {
	groups := make(map[int][]string)
	for _, v := range words {
		groups[len(v)] = append(groups[len(v)], v)
	}
	return groups
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
