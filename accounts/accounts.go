package accounts

import "fmt"

// Accounts struct shows the blueprint of Accounts
type Accounts struct {
	Number  int
	Balance float64
}

// Deposit adds to the account balance
func (user *User) Deposit(amount float64) {
	user.Acc.Balance += amount
}

// Withdraw remove from the account balance
func (user *User) Withdraw(amount float64) {
	if user.Acc.Balance-amount < 0 {
		fmt.Println("Cannot withdraw more than your account balance")
	} else {
		user.Acc.Balance -= amount
	}
}
