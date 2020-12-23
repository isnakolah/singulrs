package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ItemGet returns the newsfeeds
func ItemGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		// sugar, _ := transactions.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumia": 110}, "kg(s)")
		body := map[string]map[string]float64{
			"data": map[string]float64{
				"food": 123,
			},
		}
		c.JSON(http.StatusOK, body)
	}
}
