/*
Package slicekata covers arrays and slices: indexing, len, range loops,
returning a new slice vs mutating in place, and how append shares memory.

Every function starts as a stub that returns a zero value. Replace the body
with a real implementation until the tests go green.
*/
package slicekata

import (
	"errors"
	"slices"
)

var (
	ErrEmpty      = errors.New("empty slice")
	ErrOutOfRange = errors.New("index out of range")
)

func SumAll(nums []int) int {
	var count int
	for _, i := range nums {
		count += i
	}
	return count
}

func DoubleAll(nums []int) []int {
	result := []int{}
	if len(nums) == 0 {
		return result
	}
	for _, v := range nums {
		result = append(result, v*2)
	}
	return result
}

func FindMax(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrEmpty
	}
	return slices.Max(nums), nil;
}

func ReverseInPlace(nums []int) {
	for i, j := 0, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func CountOccurrences(words []string, target string) int {
	var count int;
	for _, i := range words {
		if target == i {
			count ++
		}
	}
	return count
}

func Flatten(nested [][]int) []int {
	var result []int
	if (len(nested) == 0) {
		return result
	}

	for _, i := range nested {
		result = append(result, i...)
	}
	return result
}

func RemoveAt(nums []int, i int) ([]int, error) {
	if i < 0 || i >= len(nums) {
		return nil, ErrOutOfRange
	}
	return slices.Delete(slices.Clone(nums), i, i+1), nil
	
}
