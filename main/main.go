package main

import (
	"bankGolang/main/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// account, err := accounts.NewUser("Test", "Account", 20, 2000)
	r := gin.Default()

	r.GET("/healthz", handler.HealthGet())
	r.GET("/item", handler.ItemGet())

	r.Run()
}
