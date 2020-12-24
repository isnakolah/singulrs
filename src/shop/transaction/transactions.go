package transaction

import (
	"errors"
	"singulr/src/shop/item"
)

// Transaction struct defines the blueprint of the transaction
type Transaction struct {
	Item      *item.Item         `json:"item"`
	ItemBrand map[string]float64 `json:"brand"`
	Amount    uint               `json:"amount"`
}

// #TODO Get item to be used in transaction

// New function creates new transaction
// If the brand is not in the list of item brands
// The function adds it to the item brands slice
func New(item *item.Item, itemBrand map[string]float64, amount uint) (transaction *Transaction, message string, err error) {
	if amount > 0 {
		transaction = &Transaction{}
		transaction.Item = item
		transaction.Amount = amount
		transaction.ItemBrand = itemBrand

		message = transaction.Item.AddBrand(itemBrand)
		return
	}
	err = errors.New("invalid amount, amount must be greater than 0")
	return
}
