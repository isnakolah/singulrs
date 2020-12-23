package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthGet returns the statuscode of OK and JSON response of a map
func HealthGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"health": "OK",
		})
	}
}
