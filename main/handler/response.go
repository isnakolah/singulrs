package handler

import (
	"bankGolang/transactions"
)

// ItemGetResponse struct to define how the GET response will look like
type ItemGetResponse struct {
	Data    Data                 `json:"data"`
	Error   Error                `json:"error"`
	Message transactions.Message `json:"message"`
}

// Data struct shapes how the data will look like
type Data struct {
	Name string `json:"name"`
}

// Error struct shapes how the error will look like
type Error struct {
	Message string `json:"message"`
}

// func GetData() {

// }
