/*
Largest uses a constraint, cmp.Ordered, so the function can use > on T.
Temperature shows a named type implementing fmt.Stringer.
*/
package ifacekata

import "cmp"

func Largest[T cmp.Ordered](xs []T) (T, bool) {
	var zero T
	return zero, false
}

type Celsius float64

func (c Celsius) String() string {
	return ""
}
