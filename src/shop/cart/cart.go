package cart

import "bankGolang/src/shop/transaction"

type cart struct {
	Items []transaction.Transaction `json:"items"`
}
