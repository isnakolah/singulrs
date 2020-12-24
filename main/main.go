package main

import (
	"bankGolang/main/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", handler.Healthz())
	r.GET("/item", itemhandler.Get())
	r.GET("/transaction", transactionhandler.Get())

	r.Run()
}
