/*
Package funckata covers functions as values: passing callbacks, closures,
your own generic Map/Filter/Reduce, and the standard slices, maps and cmp
packages that replace most of JS's array methods.
*/
package funckata

func Map[T, U any](xs []T, f func(T) U) []U {
	return nil
}

func Filter[T any](xs []T, keep func(T) bool) []T {
	return nil
}

func Reduce[T, A any](xs []T, initial A, f func(A, T) A) A {
	var zero A
	return zero
}

func Chunk[T any](xs []T, size int) [][]T {
	return nil
}

func MakeCounter() func() int {
	return func() int { return 0 }
}
