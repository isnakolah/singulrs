package cart

import "bankGolang/shop/transaction"

type cart struct {
	Items []transaction.Transaction `json:"items"`
}
