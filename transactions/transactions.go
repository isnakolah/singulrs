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

	// findBrand := func() (brandName string, brandPrice float64) {
	// 	for _, element := range transaction.Item.Brands {
	// 		if element.Name == transaction.ItemBrand.Name {
	// 			brandName = element.Name
	// 			brandPrice = element.Price
	// 		}
	// 	}
	// 	return
	// }
	// name, _ := findBrand()
	// if name == "" {
	// 	transaction.Item.AddBrand(map[string]float64{transaction.ItemBrand.Name: transaction.ItemBrand.Price})
	// }

	return
}
