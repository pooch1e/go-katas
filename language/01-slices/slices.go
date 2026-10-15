/*
Package slicekata covers arrays and slices: indexing, len, range loops,
returning a new slice vs mutating in place, and how append shares memory.

Every function starts as a stub that returns a zero value. Replace the body
with a real implementation until the tests go green.
*/
package slicekata

import "errors"

var (
	ErrEmpty      = errors.New("empty slice")
	ErrOutOfRange = errors.New("index out of range")
)

func SumAll(nums []int) int {
	return 0
}

func DoubleAll(nums []int) []int {
	return nil
}

func FindMax(nums []int) (int, error) {
	return 0, nil
}

func ReverseInPlace(nums []int) {
}

func CountOccurrences(words []string, target string) int {
	return 0
}

func Flatten(nested [][]int) []int {
	return nil
}

func RemoveAt(nums []int, i int) ([]int, error) {
	return nil, nil
}
