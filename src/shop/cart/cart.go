package cart

import "singulr/src/shop/transaction"

type cart struct {
	Items []transaction.Transaction `json:"items"`
}
