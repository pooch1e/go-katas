/*
FetchAll tests — Promise.all for Go: run every fetch concurrently, return
results in input order, fail if any fetch fails.
*/
package concurrencykata_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	concurrencykata "github.com/pooch1e/go-katas/language/12-concurrency"
)

func slowUpper(d time.Duration) func(string) (string, error) {
	return func(s string) (string, error) {
		time.Sleep(d)
		return strings.ToUpper(s), nil
	}
}

func TestFetchAll(t *testing.T) {
	t.Run("returns every result in input order", func(t *testing.T) {
		got, err := concurrencykata.FetchAll([]string{"a", "b", "c"}, slowUpper(0))
		require.NoError(t, err)
		assert.Equal(t, []string{"A", "B", "C"}, got)
	})

	// Five 50ms fetches done one after another take 250ms. Concurrently
	// they take about 50ms.
	t.Run("runs the fetches concurrently", func(t *testing.T) {
		start := time.Now()
		got, err := concurrencykata.FetchAll([]string{"a", "b", "c", "d", "e"}, slowUpper(50*time.Millisecond))
		require.NoError(t, err)
		assert.Len(t, got, 5)
		assert.Less(t, time.Since(start), 150*time.Millisecond)
	})

	t.Run("returns an error if any fetch fails", func(t *testing.T) {
		boom := errors.New("boom")
		_, err := concurrencykata.FetchAll([]string{"ok", "bad"}, func(s string) (string, error) {
			if s == "bad" {
				return "", boom
			}
			return s, nil
		})
		assert.ErrorIs(t, err, boom)
	})

	t.Run("an empty list returns an empty slice", func(t *testing.T) {
		t.Skip("todo")
	})
}
