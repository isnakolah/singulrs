package transactions

import "testing"

func TestNewItem(t *testing.T) {
	sugar := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")

	if sugar.Name != "Sugar" {
		t.Error(
			"\n-> Valid item not captured",
			"\n-> For ItemName",
			"expected", "Sugar",
			"got", sugar.Name,
		)
	}

	if len(sugar.Brands) == 2 {
		t.Error(
			"\n-> Brands of item not added.",
			"\n-> For len(Brands)",
			"expected", 2,
			"got", len(sugar.Brands),
		)
	}
}
