package transactions

import (
	"bankGolang/utils"
	"testing"
)

func TestNewItem(t *testing.T) {
	// Test on a valid new item
	sugar := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")

	for _, price := range sugar.Brands[0] {
		if price != 110 {
			t.Errorf(utils.ErrorMessage(
				"Valid item not captured", "Brand Price", float64(110), price,
			))
		}
	}

	if len(sugar.Brands) != 2 {
		t.Errorf(utils.ErrorMessage(
			"Brands of item not added.", "len(Brands)", float64(2), float64(len(sugar.Brands)),
		))
	}
}

func TestAddValidBrand(t *testing.T) {
	// Test for adding a valid brand
	sugar := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	sugar.AddBrand(map[string]float64{"Nzoia": 134})

	if len(sugar.Brands) != 3 {
		t.Errorf(utils.ErrorMessage(
			"Additional Brand not added.", "len(Brands)", float64(3), float64(len(sugar.Brands)),
		))
	}
}

func TestNewValidTransaction(t *testing.T) {
	// Test for adding a valid transaction
	sugar := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	firstPurchase := NewTransaction(sugar, map[string]float64{"Kabras": 110}, 3)

	if firstPurchase.Amount != 3 {
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

	// Test for adding a purchase with a new brand
	secondPurchase := NewTransaction(sugar, map[string]float64{"Nzoia": 200}, 1)

	if len(sugar.Brands) != 3 || len(secondPurchase.Item.Brands) != 3 {
		t.Errorf(utils.ErrorMessage(
			"None existing brand not added", "Brand length", float64(3), float64(len(secondPurchase.Item.Brands)),
		))
	}
}
