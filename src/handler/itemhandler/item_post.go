package itemhandler

import (
	"net/http"
	"singulr/src/shop/item"
	"singulr/src/utils/responseerr"

	"github.com/gin-gonic/gin"
)

func createItem(postedItem *item.Item) (createdItem *item.Item, err error) {
	createdItem, err = item.New(postedItem.Name, postedItem.Brands, postedItem.Unit)
	return
}

// PostItem function handles posting of an item
func PostItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := item.Item{}
		c.Bind(&requestBody)
		// c.Status(http.StatusNoContent)
		_, err := createItem(&requestBody)
		c.JSON(http.StatusCreated, map[string]string{"error": responseerr.GetStrErr(err)})
	}
}
