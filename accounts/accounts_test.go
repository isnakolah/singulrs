package accounts

import "testing"

func TestNew(t *testing.T) {
	jane, err := New("Jane", "Doe", 18, 1, -1234)

	if err == nil && jane.Acc.Balance != 0 {
		t.Errorf("Error was not raised for depositing a negative value")
	}

	john, err := New("John", "Doe", 17, 1, 1000)

	if err == nil && john.Name.FirstName != "" {
		t.Errorf("Error was not raised for younger age than 18")
	}

}

func TestDeposit(t *testing.T) {
	jane, _ := New("Jane", "Doe", 18, 1, 1000)

	jane.Deposit(1000)

	if jane.Acc.Balance != 2000 {
		t.Errorf("Amount not deposited")
	}
}

func TestWithdraw(t *testing.T) {
	jane, _ := New("Jane", "Doe", 18, 1, 1000)

	jane.Withdraw(2000)

	if jane.Acc.Balance != 1000 {
		t.Errorf("Withdrawn more than available in the account")
	}

	jane.Withdraw(500)

	if jane.Acc.Balance != 500 {
		t.Errorf("Withdraw was not successfull")
	}
}
