package transactions

// Transaction struct defines the blueprint of the transaction
type Transaction struct {
	Item      *Item
	ItemBrand Brand
	Amount    uint
}

// NewTransaction function creates new transaction
func NewTransaction(item *Item, itemBrand map[string]float64, amount uint) (transaction *Transaction) {
	transaction = &Transaction{}

	transaction.Item = item
	transaction.Amount = amount

	for brand, price := range itemBrand {
		transaction.ItemBrand = Brand{brand, price}
	}

	return
}
