package main

import (
	"singulr/src/handler/healthhandler"
	"singulr/src/handler/itemhandler"
	"singulr/src/handler/transactionhandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", healthhandler.Healthz())
	r.GET("/item", itemhandler.GetItem())
	r.GET("/transaction", transactionhandler.GetTransaction())

	r.Run()
}
