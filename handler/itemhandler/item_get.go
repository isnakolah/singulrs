package itemhandler

import (
	"bankGolang/shop/item"
	"bankGolang/utils/responseerr"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Data struct shapes how the data will look like
type Data struct {
	Item *item.Item `json:"item"`
}

// GetResponse struct to define how the GET response will look like
type GetResponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// GetItem returns an item
func GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		sugar, err := item.NewItem(
			"Sugar", map[string]float64{"Kabras": 110, "Mumias": 110}, "",
		)

		c.JSON(
			http.StatusOK,
			GetResponse{Data{sugar}, "", responseerr.GetStrErr(err)},
		)
	}
}
