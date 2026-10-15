/*
Stack[T] is your own generic container: a type parameter on a struct, and
methods that return (T, bool) when there may be nothing to return.
*/
package ifacekata

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	return zero, false
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	return zero, false
}

func (s *Stack[T]) Len() int {
	return 0
}
