package accounts

import (
	"errors"
)

// Accounts struct shows the blueprint of Accounts
type Accounts struct {
	Number  uint
	Balance float64
}

// Deposit adds to the account balance
func (user *User) Deposit(amount float64) {
	user.Acc.Balance += amount
}

// Withdraw remove from the account balance
func (user *User) Withdraw(amount float64) error {
	if user.Acc.Balance-amount < 0 {
		err := errors.New("invalid withdrawal, cannot withdraw amount greater than account balance")
		return err
	}
	user.Acc.Balance -= amount
	return nil
}
