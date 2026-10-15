/*
ParseAge tests — turn a string into an age between 0 and 150, wrapping the
underlying error so callers keep the detail.
*/
package errorkata_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	errorkata "github.com/pooch1e/go-katas/language/05-errors"
)

func TestParseAge(t *testing.T) {
	t.Run("parses a valid age", func(t *testing.T) {
		got, err := errorkata.ParseAge("36")
		require.NoError(t, err)
		assert.Equal(t, 36, got)
	})

	// strconv.Atoi already returns a useful error. Wrap it with
	// fmt.Errorf("...: %w", err) so the cause is still findable.
	t.Run("wraps the strconv error for non-numbers", func(t *testing.T) {
		_, err := errorkata.ParseAge("thirty")
		assert.ErrorIs(t, err, strconv.ErrSyntax)
	})

	t.Run("mentions the bad input in the message", func(t *testing.T) {
		_, err := errorkata.ParseAge("thirty")
		require.Error(t, err)
		assert.Contains(t, err.Error(), `"thirty"`)
	})

	t.Run("returns ErrAgeOutOfRange for negative ages", func(t *testing.T) {
		_, err := errorkata.ParseAge("-1")
		assert.ErrorIs(t, err, errorkata.ErrAgeOutOfRange)
	})

	t.Run("returns ErrAgeOutOfRange above 150", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("accepts the boundaries 0 and 150", func(t *testing.T) {
		t.Skip("todo")
	})
}
