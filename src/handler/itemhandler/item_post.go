package itemhandler

import (
	"singulr/src/shop/item"

	"github.com/gin-gonic/gin"
)

// type itemPostRequest struct {
// 	Name    string
// 	Brands  map[string]int
// 	Unit string
// }

// PostItem function handles posting of an item
func PostItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := item.Item{}
		c.Bind(&requestBody)
	}
}
