package main

import (
	"bankGolang/main/handler"
	"bankGolang/main/handler/itemhandler"
	"bankGolang/main/handler/transactionhandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", handler.Healthz())
	r.GET("/item", itemhandler.GetItem())
	r.GET("/transaction", transactionhandler.GetTransaction())

	r.Run()
}
