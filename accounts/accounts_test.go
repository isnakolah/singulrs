package accounts

import "testing"

func TestNewUser(t *testing.T) {
	// Test for a valid input
	daniel, err := NewUser("Daniel", "Nakolah", 20, -1000)

	if err != nil {
		t.Error(
			"\n\n-> Vaid deposit not successful",
			"For", "Account Balance",
			"expected", 1000,
			"got", daniel.Acc.Balance,
			"\n-> Error:", err, "\n",
		)
	}

	// Test for an invalid account balance
	jane, err := NewUser("Jane", "Doe", 18, -1234)

	if err == nil && jane.Acc.Balance != 0 {
		t.Error(
			"\n-> Invalid deposit, error not raised",
			"\n-> For", jane.Acc.Balance,
			"expected", 0,
			"got", jane.Acc.Balance,
			"\n-> Error:", err,
		)
	}

	// Test for an invalid age
	john, err := NewUser("John", "Doe", 17, 1000)

	if err == nil && john.Name.FirstName != "" {
		t.Error(
			"-> Invalid age, error not raised",
			"\n-> For", john.Age,
			"expected", 0,
			"got", john.Age,
			"\n-> Error:", err,
		)
	}

}

func TestDeposit(t *testing.T) {
	// Test for a valid deposit
	jane, err := NewUser("Jane", "Doe", 18, 1000)

	jane.Deposit(1000)

	if jane.Acc.Balance != 2000 {
		t.Error(
			"\n-> Valid deposit not successful",
			"\n-> For", jane.Name.FirstName, "account",
			"expected", 2000,
			"got", jane.Acc.Balance,
			"\n-> Error:", err,
		)
	}
}

func TestWithdraw(t *testing.T) {
	// Test for valid withdrawal
	john, err := NewUser("John", "Doe", 23, 1000)
	john.Withdraw(500)

	if john.Acc.Balance != 500 {
		t.Error(
			"\n-> Valid withdrawal not successful",
			"\n-> For", john.Name.FirstName, "account",
			"expected", 500,
			"got", john.Acc.Balance,
		)
	}

	// Test for invalid withdrawal
	jane, err := NewUser("Jane", "Doe", 18, 1000)

	jane.Withdraw(2000)

	if jane.Acc.Balance != 1000 && err == nil {
		t.Error(
			"\n-> Invalid withdrawal, error not raised",
			"\n-> For", jane.Name.FirstName, "account",
			"expected", 1000,
			"got", john.Acc.Balance,
			"\n-> Error:", err,
		)
	}

}
