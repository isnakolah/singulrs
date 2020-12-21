package main

import "fmt"

// Accounts struct shows the blueprint of Accounts
type Accounts struct {
	Number  int
	Balance int
}

// User is the blueprint of the user
type User struct {
	Name Names
	Age  int
	Acc  Accounts
}

// Names shows the names of the user
type Names struct {
	FirstName, MiddleName, LastName string
}

func userInit(firstName, middleName, lastName string, age, balance, accNumber int) *User {
	user := new(User)
	user.Name.FirstName, user.Name.MiddleName, user.Name.LastName = firstName, middleName, lastName
	user.Age = 20
	user.Acc.Balance, user.Acc.Number = 1000, 1
	return user
}

func main() {
	david := userInit("David", "Aniel", "Beast", 33, 3000, 23)

	fmt.Println(david.Acc)
}
