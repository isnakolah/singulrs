package account_test

import (
	"singulr/src/account"
	"singulr/src/utils/errmessages"
	"testing"
)

func TestValidNewUser(t *testing.T) {
	// Test for a valid input
	daniel, err := account.New("Daniel", "Nakolah", 20, -1000)

	if err == nil {
		t.Errorf(errmessages.ErrorMessage(
			"Valid deposit not successful", "Account Balance", float64(1000), daniel.Acc.Balance,
		))
	}
}

func TestInvalidAccountBalanceNewUser(t *testing.T) {
	// Test for an invalid account balance
	jane, err := account.New("Jane", "Doe", 18, -1234)

	if err == nil {
		t.Errorf(errmessages.ErrorMessage(
			"Invalid deposit, error not raised", "Account Balance", float64(0), jane.Acc.Balance,
		))
	}
}

func TestInvalidNameNewuser(t *testing.T) {
	// Test for an invalid name
	jane, err := account.New("", "Doe", 18, 1000)

	if err == nil || jane != nil {
		t.Errorf(errmessages.ErrorMessage(
			"Invalid name, error not raised", "Account Balance", float64(0), float64(0),
		))
	}
}

func TestInvalidAgeNewUser(t *testing.T) {
	// Test for an invalid age
	john, _ := account.New("John", "Doe", 17, 1000)

	if john != nil {
		t.Errorf(errmessages.ErrorMessage(
			"Invalid age, error not raised", "Age", float64(0), float64(john.Age),
		))
	}
}

func TestValidDeposit(t *testing.T) {
	// Test for a valid deposit
	jane, err := account.New("Jane", "Doe", 18, 1000)
	jane.Deposit(1000)

	if jane.Acc.Balance != 2000 || err != nil {
		t.Errorf(errmessages.ErrorMessage(
			"Valid deposit not successful", "Account Balance", float64(2000), jane.Acc.Balance,
		))
	}
}

func TestInvalidDeposit(t *testing.T) {
	// Test for a valid deposit
	jane, _ := account.New("Jane", "Doe", 18, 1000)
	err := jane.Deposit(0)

	if err == nil {
		t.Errorf(errmessages.ErrorMessage(
			"Invalid deposit, error not raised", "deposit", float64(1000), float64(jane.Acc.Balance),
		))
	}
}

func TestValidWithdraw(t *testing.T) {
	// Test for valid withdrawal
	john, _ := account.New("John", "Doe", 23, 1000)
	john.Withdraw(500)

	if john.Acc.Balance != 500 {
		t.Errorf(errmessages.ErrorMessage(
			"Valid withdrawal not successful", "Account Balance", float64(500), john.Acc.Balance,
		))
	}
}

func TestInvalidWithdraw(t *testing.T) {
	// Test for invalid withdrawal
	jane, err := account.New("Jane", "Doe", 18, 1000)
	jane.Withdraw(2000)

	if jane.Acc.Balance != 1000 || err != nil {
		t.Errorf(errmessages.ErrorMessage(
			"Invalid withdrawal, error not raised", "Account Balance", float64(1000), jane.Acc.Balance,
		))
	}
}
