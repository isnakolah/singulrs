package transactions

import "testing"

func TestNewItem(t *testing.T) {
	sugar := NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")

	if sugar.Name != "Sugar" {
		t.Errorf("New item not captured")
	}
}
