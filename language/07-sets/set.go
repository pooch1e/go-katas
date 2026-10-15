/*
Set is a generic set built on map[T]struct{}, the idiomatic Go set. Go has
no built-in set type; this is what everyone writes.
*/
package setkata

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	return &Set[T]{}
}

func (s *Set[T]) Add(item T) bool {
	return false
}

func (s *Set[T]) Has(item T) bool {
	return false
}

func (s *Set[T]) Remove(item T) {
}

func (s *Set[T]) Len() int {
	return 0
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	return &Set[T]{}
}
