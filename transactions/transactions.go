package transactions

// Transaction struct defines the blueprint of the transaction
type Transaction struct {
	Item   Item
	Amount uint
}

// NewTransaction function creates new transaction
func NewTransaction(item Item, amount uint) (transaction *Transaction) {
	transaction = &Transaction{}

	transaction.Item = item
	transaction.Amount = amount
	return
}
