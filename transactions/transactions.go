package transactions

// Transaction struct defines the blueprint of the transaction
type Transaction struct {
	Item      *Item
	ItemBrand map[string]float64
	Amount    uint
}

// NewTransaction function creates new transaction
func NewTransaction(item *Item, itemBrand map[string]float64, amount uint) (transaction *Transaction) {
	transaction = &Transaction{}

	transaction.Item = item
	transaction.Amount = amount
	transaction.ItemBrand = itemBrand

	var brandName string
	for name, _ := range itemBrand {
		brandName = name
	}
	findBrand := func() (available bool) {
		available = false
		for _, item := range transaction.Item.Brands {
			if _, ok := item[brandName]; ok {
				available = true
			}
		}
		return
	}
	if !findBrand() {
		transaction.Item.AddBrand(itemBrand)
	}

	return
}
