package accounts

import (
	"errors"
)

// User is the blueprint of the user
type User struct {
	Name Names
	Age  uint
	Acc  Accounts
}

// Names shows the names of the user
type Names struct {
	FirstName, MiddleName, LastName string
}

// New to create a new User
func New(firstName, lastName string, age uint, balance float64) (*User, error) {
	if age < 18 {
		err := errors.New("invalid age, under 18 years")
		return &User{}, err
	}
	user := &User{}
	user.Name.FirstName, user.Name.LastName = firstName, lastName
	user.Age, user.Acc.Number = age, accNumber()

	if balance >= 0 {
		user.Acc.Balance = balance
	} else {
		err := errors.New("invalid account balance, should be greater than 0")
		user.Acc.Balance = 0
		return user, err
	}
	return user, nil
}
