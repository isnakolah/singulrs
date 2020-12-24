package handler

import (
	"bankGolang/transactions"

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

	}
}
