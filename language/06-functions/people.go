/*
People queries — the Go answer to LINQ / lodash chains over a slice of
structs, using the slices, maps and cmp packages.
*/
package funckata

type Person struct {
	Name string
	Age  int
	City string
}

func SortedByAgeThenName(people []Person) []Person {
	return nil
}

func Oldest(people []Person) (Person, bool) {
	return Person{}, false
}

func CountByCity(people []Person) map[string]int {
	return nil
}

func AverageAge(people []Person) (float64, bool) {
	return 0, false
}
