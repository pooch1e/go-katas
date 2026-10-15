/*
Point is a small value type: a struct with value-receiver methods that
return new Points instead of changing the one they were called on.
*/
package structkata

type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	return Point{}
}

func (p Point) Moved(dx, dy int) Point {
	return Point{}
}

func (p Point) DistanceTo(other Point) float64 {
	return 0
}
