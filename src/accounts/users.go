package accounts

import (
	"bankGolang/src/utils/checkstring"
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

// AddDetails methods adds the details to the user
func (user *User) AddDetails(firstName, lastName string, age uint) {
	user.Name.FirstName = firstName
	user.Name.LastName = lastName
	user.Age = age
	user.Acc.Number = AccNumber()
}

// NewUser to create a new User
func NewUser(firstName, lastName string, age uint, balance float64) (user *User, err error) {
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
