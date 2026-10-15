/*
Shared test data for the people katas.
*/
package funckata_test

import funckata "github.com/pooch1e/go-katas/language/06-functions"

func people() []funckata.Person {
	return []funckata.Person{
		{Name: "Ada", Age: 36, City: "London"},
		{Name: "Linus", Age: 54, City: "Portland"},
		{Name: "Grace", Age: 85, City: "New York"},
		{Name: "Alan", Age: 41, City: "London"},
		{Name: "Barbara", Age: 36, City: "Boston"},
	}
}
