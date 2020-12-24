package cart

import "singulr/src/shop/transaction"

// Cart struct is a blueprint for the cart
type Cart struct {
	Items []*transaction.Transaction `json:"items"`
}

// Add method adds items to Items list
func (cart *Cart) Add(items []*transaction.Transaction) {
	if items == nil {
		cart.Items = append(cart.Items, &transaction.Transaction{})
	}
	for _, item := range items {
		cart.Items = append(cart.Items, item)
	}
}

// New creates a new transaction
func New(items ...*transaction.Transaction) (cart Cart) {
	cart.Add(items)
	return
}
