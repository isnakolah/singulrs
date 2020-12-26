package account

import (
	"errors"
	"singulr/src/bank"
	"singulr/src/utils/checkstring"
)

// #TODO Use DOB to get age

// User is the blueprint of the user
type User struct {
	Name Names
	Age  uint
	Acc  bank.Accounts
}

// Names shows the names of the user
type Names struct {
	FirstName, MiddleName, LastName string
}

// New to create a new User
func New(firstName, lastName string, age uint, balance float64) (user *User, err error) {
	if !checkstring.CheckString(firstName) || !checkstring.CheckString(lastName) {
		err = errors.New("invalid name, cannot be a blank string")
		return
	} else if age < 18 {
		err = errors.New("invalid age, should be 18 years or older")
		return
	}

	user = &User{}
	user.AddDetails(firstName, lastName, age)

	if balance >= 0 {
		user.Acc.Balance = balance
	} else {
		err = errors.New("invalid account balance, should be greater than 0")
		user.Acc.Balance = 0
	}
	return
}
