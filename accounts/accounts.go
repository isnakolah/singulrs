package accounts

import "fmt"

// Accounts struct shows the blueprint of Accounts
type Accounts struct {
	Number  int
	Balance float64
}

func (user *User) deposit(amount float64) {
	user.Acc.Balance += amount
}

func (user *User) withdraw(amount float64) {
	if user.Acc.Balance-amount < 0 {
		fmt.Println("Cannot withdraw more than your account balance")
	} else {
		user.Acc.Balance -= amount
	}
}
