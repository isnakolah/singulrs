package handler

import (
	"bankGolang/transactions"
)

// Data struct shapes how the data will look like
type Data struct {
	// Item map[string]map[string]float64 `json:"item"`
	Item *transactions.Item `json:"item"`
}

// ItemGetResponse struct to define how the GET response will look like
type ItemGetResponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// GetData function returns the data
func GetData(item *transactions.Item, message, err string) (response ItemGetResponse) {
	response.Data.Item = item
	response.Message = message
	response.Error = err
	return
}
