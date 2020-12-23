package transactions

import (
	"errors"
	"fmt"
)

// Transaction struct defines the blueprint of the transaction
type Transaction struct {
	Item      *Item
	ItemBrand map[string]float64
	Amount    uint
}

// #TODO Get item to be used in transaction

// NewTransaction function creates new transaction
// If the brand is not in the list of item brands
// The function adds it to the item brands slice
func NewTransaction(item *Item, itemBrand map[string]float64, amount uint) (transaction *Transaction, message string, err error) {
	if amount > 0 {
		transaction = &Transaction{}
		transaction.Item = item
		transaction.Amount = amount
		transaction.ItemBrand = itemBrand

		var brandName string
		for name := range itemBrand {
			brandName = name
		}
		findBrand := func() (available bool) {
			for _, item := range transaction.Item.Brands {
				if _, ok := item[brandName]; ok {
					available = true
				}
			}
			return
		}
		if !findBrand() {
			transaction.Item.AddBrand(itemBrand)
			message = fmt.Sprintf("%s brand created.", brandName)
		}
		return
	}
	err = errors.New("invalid amount, amount must be greater than 0")
	return
}
