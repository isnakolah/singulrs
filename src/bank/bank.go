package bank

import (
	"errors"
	"math/rand"
	"singulr/src/account"
)

// Accounts struct shows the blueprint of Accounts
type Accounts struct {
	Number  int
	Balance float64
}

// Deposit adds to the account balance
func (user *account.User) Deposit(amount float64) (err error) {
	if amount > 0 {
		user.Acc.Balance += amount
		return
	}
	err = errors.New("invalid deposit, must be greater than 0")
	return
}

// Withdraw function deducts from your account
// Only when the balance is more than the value to be withdrawn
func (user *account.User) Withdraw(amount float64) (err error) {
	if user.Acc.Balance-amount < 0 {
		err = errors.New("invalid withdrawal, cannot withdraw amount greater than account balance")
		return
	}
	user.Acc.Balance -= amount
	return
}

// AddDetails methods adds the details to the user
func (user *account.User) AddDetails(firstName, lastName string, age uint) {
	user.Name.FirstName = firstName
	user.Name.LastName = lastName
	user.Age = age
	user.Acc.Number = bank.AccNumber()
}

// AccNumber function creates a random number as the account number
func AccNumber() (accNumber int) {
	accNumber = rand.Int()
	return
}
