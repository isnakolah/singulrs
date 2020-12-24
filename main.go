package main

import (
	"bankGolang/handler/healthhandler"
	"bankGolang/handler/itemhandler"
	"bankGolang/handler/transactionhandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", healthhandler.Healthz())
	r.GET("/item", itemhandler.GetItem())
	r.GET("/transaction", transactionhandler.GetTransaction())

	r.Run()
}
