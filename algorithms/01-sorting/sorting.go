/*
Package sortkata implements four classic sorts by hand. In real code you'd
call slices.Sort; here the point is knowing what it costs and why.
*/
package sortkata

import "cmp"

func BubbleSort[T cmp.Ordered](xs []T) {
}

func InsertionSort[T cmp.Ordered](xs []T) {
}

func MergeSort[T cmp.Ordered](xs []T) []T {
	return nil
}

func QuickSort[T cmp.Ordered](xs []T) {
}
