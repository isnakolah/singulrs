package accounts

import (
	"errors"
	"math/rand"
)

// Accounts struct shows the blueprint of Accounts
type Accounts struct {
	Number  int
	Balance float64
}

// Deposit adds to the account balance
func (user *User) Deposit(amount float64) {
	user.Acc.Balance += amount
}

// Withdraw function deducts from your account
// Only when the balance is more than the value to be withdrawn
func (user *User) Withdraw(amount float64) error {
	if user.Acc.Balance-amount < 0 {
		err := errors.New("invalid withdrawal, cannot withdraw amount greater than account balance")
		return err
	}
	user.Acc.Balance -= amount
	return nil
}

// AccNumber function creates a random number as the account number
func AccNumber() (accNumber int) {
	accNumber = rand.Int()
	return
}
