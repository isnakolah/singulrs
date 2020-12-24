package cart_test

import (
	"singulr/src/shop/cart"
	"singulr/src/shop/item"
	"singulr/src/shop/transaction"
	"singulr/src/utils/errmessages"
	"testing"
)

func TestNewCart(t *testing.T) {
	sugar, _ := item.New("Sugar", map[string]float64{"Kabras": 110}, "kg(s)")
	rice, _ := item.New("Rice", map[string]float64{"Pishori": 150}, "kg(s)")

	purchase1, _, _ := transaction.New(sugar, map[string]float64{"Kabras": 110}, 2)
	purchase2, _, _ := transaction.New(rice, map[string]float64{"Pishori": 150}, 2)

	cart := cart.New(purchase1, purchase2)

	if len(cart.Items) != 2 {
		t.Errorf(errmessages.ErrorMessage(
			"Cart items not added", "len(cart.Items)", float64(2), float64(len(cart.Items)),
		))
	}
}
