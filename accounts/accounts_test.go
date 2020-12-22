package accounts

import "testing"

func TestNewUser(t *testing.T) {
	// Test for a valid input
	daniel, err := NewUser("Daniel", "Nakolah", 20, 1000)

	if err != nil && daniel.Name.FirstName != "Daniel" {
		t.Errorf("A valid input not successfully created")
	}

	// Test for an invalid account balance
	jane, err := NewUser("Jane", "Doe", 18, -1234)

	if err == nil && jane.Acc.Balance != 0 {
		t.Errorf("Error not raised for depositing a negative value")
	}

	// Test for an invalid age
	john, err := NewUser("John", "Doe", 17, 1000)

	if err == nil && john.Name.FirstName != "" {
		t.Errorf("Error not raised for younger age than 18")
	}

}

func TestDeposit(t *testing.T) {
	// Test for a valid deposit
	jane, _ := NewUser("Jane", "Doe", 18, 1000)

	jane.Deposit(1000)

	if jane.Acc.Balance != 2000 {
		t.Errorf("Amount not successfully deposited")
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
		t.Errorf("Error not raised for withdrawal of more than the available in the account")
	}

}
