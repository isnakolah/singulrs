package transactionhandler

import (
	"bankGolang/shop/item"
	"bankGolang/shop/transaction"
	"bankGolang/utils/responseerr"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetData is a blueprint of how the transaction will be
type GetData struct {
	Transaction *transaction.Transaction `json:"transaction"`
}

// GetResponse is blueprint for the esponse
type GetResponse struct {
	Data    GetData `json:"data"`
	Message string  `json:"message"`
	Error   string  `json:"error"`
}

// GetTransaction returns a transaction
func GetTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		sugar, _ := item.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 110}, "kg(s)")
		purchase, message, err := transaction.NewTransaction(sugar, map[string]float64{"Nzoia": 150}, 3)
		c.JSON(
			http.StatusOK,
			GetResponse{GetData{purchase}, message, responseerr.GetStrErr(err)},
		)
	}
}
