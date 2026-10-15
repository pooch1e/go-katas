/*
ValidateUser tests — a custom error type carrying structured detail, pulled
back out with errors.As, and errors.Join for reporting several problems.

Rules: Name must not be blank, Email must contain "@".
*/
package errorkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	errorkata "github.com/pooch1e/go-katas/language/05-errors"
)

func TestValidateUser(t *testing.T) {
	t.Run("a valid user returns nil", func(t *testing.T) {
		assert.NoError(t, errorkata.ValidateUser(errorkata.User{Name: "Ada", Email: "ada@example.com"}))
	})

	t.Run("ValidationError formats as field: reason", func(t *testing.T) {
		err := &errorkata.ValidationError{Field: "email", Reason: "must contain @"}
		assert.Equal(t, "email: must contain @", err.Error())
	})

	// errors.As is the typed version of errors.Is: "is there a
	// *ValidationError anywhere in this chain? If so, put it in ve."
	t.Run("a bad email gives a ValidationError for the email field", func(t *testing.T) {
		err := errorkata.ValidateUser(errorkata.User{Name: "Ada", Email: "nope"})
		var ve *errorkata.ValidationError
		require.ErrorAs(t, err, &ve)
		assert.Equal(t, "email", ve.Field)
	})

	t.Run("a blank name gives a ValidationError for the name field", func(t *testing.T) {
		t.Skip("todo")
	})

	// errors.Join(errA, errB) returns one error wrapping both. errors.As
	// and errors.Is search every branch. How would you check that BOTH
	// fields were reported?
	t.Run("reports every failing field, not just the first", func(t *testing.T) {
		t.Skip("todo")
	})
}
