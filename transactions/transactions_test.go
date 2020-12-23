package transactions

import (
	"bankGolang/utils"
	"testing"
)

func TestNewItem(t *testing.T) {
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
