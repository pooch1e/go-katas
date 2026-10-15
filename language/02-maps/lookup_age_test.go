/*
LookupAge tests — return a person's age, or the fallback when they're missing.
*/
package mapkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mapkata "github.com/pooch1e/go-katas/language/02-maps"
)

func TestLookupAge(t *testing.T) {
	ages := map[string]int{"ada": 36, "linus": 54, "newborn": 0}

	t.Run("returns the age when the name exists", func(t *testing.T) {
		assert.Equal(t, 36, mapkata.LookupAge(ages, "ada", -1))
	})

	t.Run("returns the fallback when the name is missing", func(t *testing.T) {
		assert.Equal(t, -1, mapkata.LookupAge(ages, "grace", -1))
	})

	// ages["grace"] gives 0 for a missing key, and so does ages["newborn"].
	// Checking "if age == 0" can't tell those apart. What can?
	t.Run("returns a stored zero rather than the fallback", func(t *testing.T) {
		assert.Equal(t, 0, mapkata.LookupAge(ages, "newborn", -1))
	})

	t.Run("works on a nil map", func(t *testing.T) {
		assert.Equal(t, 99, mapkata.LookupAge(nil, "anyone", 99))
	})
}
