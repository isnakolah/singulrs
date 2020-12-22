package accounts

import "testing"

func TestNewUser(t *testing.T) {
	// Test for a valid input
	daniel, err := NewUser("Daniel", "Nakolah", 20, 1000)

	if err != nil && daniel.Name.FirstName != "Daniel" {
		t.Error(
			"-> Vaid deposit not successful\n\t\t",
			"-> For", daniel.Acc.Balance,
			"expected", 1000,
			"got", daniel.Acc.Balance, "\n\t\t",
			"-> Error:", err,
		)
		t.Errorf("A valid input not successfully created")
	}

	// Test for an invalid account balance
	jane, err := NewUser("Jane", "Doe", 18, -1234)

	if err == nil && jane.Acc.Balance != 0 {
		t.Error(
			"-> Invalid deposit, error not raised\n\t\t",
			"-> For", jane.Acc.Balance,
			"expected", 0,
			"got", jane.Acc.Balance, "\n\t\t",
			"-> Error:", err,
		)
	}

	// Test for an invalid age
	john, err := NewUser("John", "Doe", 17, 1000)

	if err == nil && john.Name.FirstName != "" {
		t.Error(
			"-> Invalid age, error not raised\n\t\t",
			"-> For", john.Age,
			"expected", 0,
			"got", john.Age, "\n\t\t",
			"-> Error:", err,
		)
	}

}

func TestDeposit(t *testing.T) {
	// Test for a valid deposit
	jane, err := NewUser("Jane", "Doe", 18, 1000)

	jane.Deposit(100)

	if jane.Acc.Balance != 2000 {
		t.Error(
			"-> Valid deposit not successful\n\t\t",
			"-> For", jane.Name.FirstName, "account",
			"expected", 2000,
			"got", jane.Acc.Balance, "\n\t\t",
			"-> Error:", err,
		)
	}
}

func TestWithdraw(t *testing.T) {
	// Test for valid withdrawal
	john, err := NewUser("John", "Doe", 23, 1000)
	john.Withdraw(500)

	if john.Acc.Balance != 500 {
		t.Error(
			"-> Valid withdrawal not successful\n\t\t",
			"-> For", john.Name.FirstName, "account",
			"expected", 500,
			"got", john.Acc.Balance,
		)
	}

	// Test for invalid withdrawal
	jane, err := NewUser("Jane", "Doe", 18, 1000)

	jane.Withdraw(2000)

	if jane.Acc.Balance != 1000 && err == nil {
		t.Error(
			"-> Invalid withdrawal, error not raised\n\t\t",
			"-> For", jane.Name.FirstName, "account",
			"expected", 1000,
			"got", john.Acc.Balance, "\n\t\t",
			"-> Error:", err,
		)
	}

}
