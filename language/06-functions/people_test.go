/*
People query tests — sorting, max-by, grouping and averaging with the
standard library.
*/
package funckata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	funckata "github.com/pooch1e/go-katas/language/06-functions"
)

func TestSortedByAgeThenName(t *testing.T) {
	t.Run("sorts by age, breaking ties by name", func(t *testing.T) {
		got := funckata.Map(funckata.SortedByAgeThenName(people()), func(p funckata.Person) string { return p.Name })
		assert.Equal(t, []string{"Ada", "Barbara", "Alan", "Linus", "Grace"}, got)
	})

	// slices.SortFunc sorts in place. Sorting the caller's slice would be a
	// nasty surprise, so clone first.
	t.Run("does not reorder the input", func(t *testing.T) {
		input := people()
		funckata.SortedByAgeThenName(input)
		assert.Equal(t, people(), input)
	})
}

func TestOldest(t *testing.T) {
	t.Run("returns the oldest person", func(t *testing.T) {
		got, ok := funckata.Oldest(people())
		require.True(t, ok)
		assert.Equal(t, "Grace", got.Name)
	})

	t.Run("reports false for no people", func(t *testing.T) {
		_, ok := funckata.Oldest(nil)
		assert.False(t, ok)
	})
}

func TestCountByCity(t *testing.T) {
	t.Run("counts people per city", func(t *testing.T) {
		want := map[string]int{"London": 2, "Portland": 1, "New York": 1, "Boston": 1}
		assert.Equal(t, want, funckata.CountByCity(people()))
	})
}

func TestAverageAge(t *testing.T) {
	t.Run("averages everyone's age", func(t *testing.T) {
		t.Skip("todo")
	})

	// Integer division truncates: 7 / 2 is 3. Convert before dividing.
	t.Run("keeps the fractional part", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("reports false for no people", func(t *testing.T) {
		_, ok := funckata.AverageAge(nil)
		assert.False(t, ok)
	})
}
