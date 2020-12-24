package handler

import (
	"bankGolang/transactions"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ItemData struct shapes how the data will look like
type ItemData struct {
	// Item map[string]map[string]float64 `json:"item"`
	Item *transactions.Item `json:"item"`
}

// ItemGetResponse struct to define how the GET response will look like
type ItemGetResponse struct {
	Data    ItemData `json:"data"`
	Message string   `json:"message"`
	Error   string   `json:"error"`
}

// ItemGetData function returns the data
func ItemGetData(item *transactions.Item, message, err string) (response ItemGetResponse) {
	response.Data.Item = item
	response.Message = message
	response.Error = err
	return
}

// ItemGet returns an item
func ItemGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		sugar, err := transactions.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 110}, "")

		c.JSON(
			http.StatusOK,
			ItemGetData(sugar, "", err.Error()),
		)
	}
}
