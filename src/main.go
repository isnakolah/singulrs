package main

import (
	"bankGolang/src/handler/healthhandler"
	"bankGolang/src/handler/itemhandler"
	"bankGolang/src/handler/transactionhandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", healthhandler.Healthz())
	r.GET("/item", itemhandler.GetItem())
	r.GET("/transaction", transactionhandler.GetTransaction())

	r.Run()
}
