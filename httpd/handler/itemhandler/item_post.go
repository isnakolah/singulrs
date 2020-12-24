package itemhandler

import (
	"bankGolang/transactions"

	"github.com/gin-gonic/gin"
)

// type itemPostRequest struct {
// 	Name    string
// 	Brands  map[string]int
// 	Measure string
// }

// PostItem function handles posting of an item
func PostItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := transactions.Item{}
        c.Bind(&requestBody)
	}
}
