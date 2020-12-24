package cart

import (
	"errors"
	"singulr/src/shop/transaction"
)

// Cart struct is a blueprint for the cart
type Cart struct {
	Items []*transaction.Transaction `json:"items"`
}

// Add method adds items to Items list
func (cart *Cart) Add(items []*transaction.Transaction) (err error) {
	if items == nil {
		cart.Items = nil
		err = errors.New("no new items entered")
		return
	}
	for _, item := range items {
		cart.Items = append(cart.Items, item)
	}
	return
}

// New creates a new transaction
func New(items ...*transaction.Transaction) (cart Cart, err error) {
	err = cart.Add(items)
	return
}
