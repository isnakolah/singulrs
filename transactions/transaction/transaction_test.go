package transaction

import (
	"bankGolang/transactions/item"
	"bankGolang/utils"
	"testing"
)

func TestValidNewTransaction(t *testing.T) {
	// Test for adding a valid transaction
	sugar, _ := item.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	firstPurchase, message, err := NewTransaction(sugar, map[string]float64{"Kabras": 110}, 3)

	if firstPurchase.Amount != 3 || err != nil || message != "" {
		t.Errorf(utils.ErrorMessage(
			"Valid Transaction not successful", "Amount", float64(3), float64(firstPurchase.Amount),
		))
	}

	// Test for number of brands in the sugar item
	// NOTE: No new item created
	if len(sugar.Brands) != 2 || len(firstPurchase.Item.Brands) != 2 {
		t.Errorf(utils.ErrorMessage(
			"Invalid number of brands", "Number of Brands", float64(2), float64(len(sugar.Brands)),
		))
	}
}

func TestValidTransactionAutoAddBrand(t *testing.T) {
	// Test for a valid transaction whose brand isn't in the list of brands of the item
	sugar, _ := item.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	purchase, message, err := NewTransaction(sugar, map[string]float64{"Nzoia": 200}, 1)

	if len(sugar.Brands) != 3 || len(purchase.Item.Brands) != 3 || message == "" || err != nil {
		t.Errorf(utils.ErrorMessage(
			"None existing brand not added", "Brand length", float64(3), float64(len(purchase.Item.Brands)),
		))
	}

}

func TestInvalidAmountNewTransaction(t *testing.T) {
	// Test for an invalid transaction where the amount is 0
	sugar, _ := item.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	firstPurchase, _, _ := NewTransaction(sugar, map[string]float64{"Kabras": 110}, 0)

	if firstPurchase != nil {
		t.Errorf(utils.SimpleErrorMessage("Invalid purchase not caught, error not raised"))
	}
}
