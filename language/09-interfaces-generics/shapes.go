/*
Shapes — an interface satisfied implicitly by two struct types, a function
that works on any Shape, and a type switch for when you need the concrete
type back.
*/
package ifacekata

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Compile-time proof that both types satisfy Shape. If a method is missing
// or has the wrong signature, the build fails here rather than in a test.
var (
	_ Shape = Circle{}
	_ Shape = Rectangle{}
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}

func (c Circle) Perimeter() float64 {
	return 0
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

func (r Rectangle) Perimeter() float64 {
	return 0
}

func TotalArea(shapes []Shape) float64 {
	return 0
}

func Describe(s Shape) string {
	return ""
}
