package transactions

import (
	"bankGolang/utils"
	"testing"
)

func TestValidNewItem(t *testing.T) {
	// Test on a valid new item
	sugar, _ := NewItem("Sugar", map[string]float64{"Kabras": 110}, "kg(s)")

	for _, price := range sugar.Brands[0] {
		if price != 110 {
			t.Errorf(utils.ErrorMessage(
				"Valid item not captured", "Brand Price", float64(110), price,
			))
		}
	}

	if len(sugar.Brands) != 1 {
		t.Errorf(utils.ErrorMessage(
			"Brands of item not added.", "len(Brands)", float64(2), float64(len(sugar.Brands)),
		))
	}
}

func TestInvalidNameNewItem(t *testing.T) {
	// Test for an invalid name, empty string. Should return nil as the item and an err message
	sugar, err := NewItem("", map[string]float64{"Kabras": 110}, "kg(s)")

	if sugar != nil || err == nil {
		if sugar == nil {
			t.Errorf(utils.SimpleErrorMessage("Item is nil"))
		} else {
			t.Errorf(utils.ErrorMessage(
				"Invalid name, error not raised", "Item", float64(0), float64(sugar.Brands[0]["Kabras"]),
			))
		}
	}
}

func TestInvalidUnitNewItem(t *testing.T) {
	// Test for an invalid unit name, empty string, should return an error
	sugar, err := NewItem("sugar", map[string]float64{"Kabras": 110}, "")

	if sugar.Measure != "" || err == nil {
		t.Errorf(utils.SimpleErrorMessage("Invalid unit, error not raised"))
	}
}

func TestValidAddBrand(t *testing.T) {
	// Test for adding a valid brand
	sugar, _ := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	sugar.AddBrand(map[string]float64{"Nzoia": 134})

	if len(sugar.Brands) != 3 {
		t.Errorf(utils.ErrorMessage(
			"Additional Brand not added.", "len(Brands)", float64(3), float64(len(sugar.Brands)),
		))
	}
}

func TestValidNewTransaction(t *testing.T) {
	// Test for adding a valid transaction
	sugar, _ := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
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
	sugar, _ := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	purchase, message, err := NewTransaction(sugar, map[string]float64{"Nzoia": 200}, 1)

	if len(sugar.Brands) != 3 || len(purchase.Item.Brands) != 3 || message == "" || err != nil {
		t.Errorf(utils.ErrorMessage(
			"None existing brand not added", "Brand length", float64(3), float64(len(purchase.Item.Brands)),
		))
	}

}

func TestInvalidAmountNewTransaction(t *testing.T) {
	// Test for an invalid transaction where the amount is 0
	sugar, _ := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	firstPurchase, _, _ := NewTransaction(sugar, map[string]float64{"Kabras": 110}, 0)

	if firstPurchase != nil {
		t.Errorf(utils.SimpleErrorMessage("Invalid purchase not caught, error not raised"))
	}
}
