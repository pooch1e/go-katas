/*
BankAccount is a stateful type: unexported fields, a constructor function,
pointer-receiver methods that mutate, and errors for invalid operations.

Money is stored as whole pence in an int. Go has no decimal type, and
float64 can't represent 0.10 exactly.
*/
package structkata

import "errors"

var (
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type BankAccount struct {
	owner   string
	balance int
}

func NewBankAccount(owner string) *BankAccount {
	return &BankAccount{}
}

func (a *BankAccount) Owner() string {
	return ""
}

func (a *BankAccount) Balance() int {
	return 0
}

func (a *BankAccount) Deposit(pence int) error {
	return nil
}

func (a *BankAccount) Withdraw(pence int) error {
	return nil
}

func (a *BankAccount) TransferTo(to *BankAccount, pence int) error {
	return nil
}
