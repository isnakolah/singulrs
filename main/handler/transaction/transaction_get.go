package handler

import (
	"bankGolang/transactions"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TransactionGetData is a blueprint of how the transaction will be
type TransactionGetData struct {
	Transaction *transactions.Transaction `json:"transaction"`
}

// TransactionGetResponse is blueprint for the response
type TransactionGetResponse struct {
	Data    TransactionGetData `json:"data"`
	Message string             `json:"message"`
	Error   string             `json:"error"`
}

// TransactionGet returns a transaction
func TransactionGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		sugar, _ := transactions.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 110}, "kg(s)")
		purchase, message, err := transactions.NewTransaction(sugar, map[string]float64{"Nzoia": 150}, 3)
		c.JSON(
			http.StatusOK,
			TransactionGetResponse{TransactionGetData{purchase}, message, err.Error()},
		)
	}
}
