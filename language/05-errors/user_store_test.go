/*
UserStore tests — sentinel errors travelling up through several layers of
wrapping, each layer adding context.
*/
package errorkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	errorkata "github.com/pooch1e/go-katas/language/05-errors"
)

func TestUserStore(t *testing.T) {
	store := errorkata.NewUserStore(
		errorkata.User{ID: 1, Name: "Ada", Email: "ada@example.com"},
		errorkata.User{ID: 2, Name: "Linus", Email: "linus@example.com"},
	)

	t.Run("FindUser returns a stored user", func(t *testing.T) {
		u, err := store.FindUser(2)
		require.NoError(t, err)
		assert.Equal(t, "Linus", u.Name)
	})

	t.Run("FindUser returns ErrNotFound for an unknown id", func(t *testing.T) {
		_, err := store.FindUser(99)
		assert.ErrorIs(t, err, errorkata.ErrNotFound)
	})

	t.Run("EmailFor returns the user's email", func(t *testing.T) {
		got, err := store.EmailFor(1)
		require.NoError(t, err)
		assert.Equal(t, "ada@example.com", got)
	})

	// Each layer adds what it was doing, so the final message reads like a
	// stack trace: "email for user 99: find user 99: not found".
	t.Run("EmailFor wraps FindUser's error with context", func(t *testing.T) {
		_, err := store.EmailFor(99)
		require.ErrorIs(t, err, errorkata.ErrNotFound)
		assert.Equal(t, "email for user 99: find user 99: not found", err.Error())
	})

	t.Run("NewUserStore with no users finds nothing", func(t *testing.T) {
		t.Skip("todo")
	})
}
