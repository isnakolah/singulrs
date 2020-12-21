package accounts

import (
	"fmt"
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
func New(firstName, lastName string, age uint, accNumber uint, balance float64) *User {
	if age < 18 {
		fmt.Println("You cannot create an account when you are underage")
		return &User{}
	}
	user := &User{}
	user.Name.FirstName, user.Name.LastName = firstName, lastName
	user.Age, user.Acc.Number = age, accNumber

	if balance >= 0 {
		user.Acc.Balance = balance
	} else {
		fmt.Println("Input a valid account balance. Must be greater than Ksh 0")
		user.Acc.Balance = 0
	}
	return user
}
