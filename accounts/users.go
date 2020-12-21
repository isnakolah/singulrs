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

// New to create a new User
func New(firstName, lastName string, age, accNumber int, balance float64) *User {
	user := &User{}
	user.Name.FirstName, user.Name.LastName = firstName, lastName
	user.Acc.Balance, user.Age, user.Acc.Number = balance, age, accNumber

	return user
}
