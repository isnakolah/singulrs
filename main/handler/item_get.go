package handler

import (
	"bankGolang/transactions"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ItemGet returns the newsfeeds
func ItemGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		sugar, err := transactions.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 110}, "")

		c.JSON(
			http.StatusOK,
			GetData(sugar, "retrieval successful", err.Error()),
		)
	}
}
