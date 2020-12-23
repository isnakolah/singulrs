package handler

import (
	"bankGolang/transactions"
)

// ItemGetResponse struct to define how the GET response will look like
type ItemGetResponse struct {
	Data    Data
	Error   Error
	Message transactions.Message
}

// Data struct shapes how the data will look like
type Data struct {
	Name string
}

// Error struct shapes how the error will look like
type Error struct {
	Message string
}
