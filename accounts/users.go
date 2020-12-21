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

// func (user *User)

// NewUser to create a new User
func NewUser(firstName, lastName string, age uint, balance float64) (user *User, err error) {
	user, err = &User{}, nil

	if age < 18 {
		err = errors.New("invalid age, should be 18 years or older")
		return
	}
	user.Name.FirstName, user.Name.LastName = firstName, lastName
	user.Age, user.Acc.Number = age, accNumber()

	if balance >= 0 {
		user.Acc.Balance = balance
	} else {
		err = errors.New("invalid account balance, should be greater than 0")
		user.Acc.Balance = 0
	}
	return
}
