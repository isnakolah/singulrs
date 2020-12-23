package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello func displays hello world
func Hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<!DOCTYPE html>
        <html>
            <head>
                <title>Hello World</title>
            </head>
            <body>
                Hello, World!
            </body>
        </html>`,
	)
}

// PingGet returns the statuscode of OK and JSON response of a map
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Found me",
		})
	}
}
