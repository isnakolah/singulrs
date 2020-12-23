package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingGet returns the statuscode of OK and JSON response of a map
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
