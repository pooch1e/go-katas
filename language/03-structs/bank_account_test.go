/*
BankAccount tests — a constructor, pointer receivers that mutate, and
errors for operations that would break the account's rules.
*/
package structkata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	structkata "github.com/pooch1e/go-katas/language/03-structs"
)

func TestBankAccount(t *testing.T) {
	t.Run("a new account has an owner and a zero balance", func(t *testing.T) {
		acc := structkata.NewBankAccount("ada")
		assert.Equal(t, "ada", acc.Owner())
		assert.Equal(t, 0, acc.Balance())
	})

	t.Run("Deposit increases the balance", func(t *testing.T) {
		acc := structkata.NewBankAccount("ada")
		require.NoError(t, acc.Deposit(500))
		assert.Equal(t, 500, acc.Balance())
	})

	t.Run("Deposit rejects zero and negative amounts", func(t *testing.T) {
		acc := structkata.NewBankAccount("ada")
		assert.ErrorIs(t, acc.Deposit(0), structkata.ErrInvalidAmount)
		assert.ErrorIs(t, acc.Deposit(-10), structkata.ErrInvalidAmount)
	})

	t.Run("Withdraw decreases the balance", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("Withdraw returns ErrInsufficientFunds and leaves the balance alone", func(t *testing.T) {
		t.Skip("todo")
	})

	t.Run("TransferTo moves money between two accounts", func(t *testing.T) {
		t.Skip("todo")
	})

	// If the withdraw half fails, the deposit half must not happen.
	t.Run("a failed TransferTo changes neither account", func(t *testing.T) {
		t.Skip("todo")
	})
}
