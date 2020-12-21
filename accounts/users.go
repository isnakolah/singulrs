package accounts

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

func (user *User) construct(firstName, lastName string, age, accNumber int, balance float64) {
	user.Name.FirstName, user.Name.LastName = firstName, lastName
	user.Age = age
	user.Acc.Balance, user.Acc.Number = balance, accNumber
}
