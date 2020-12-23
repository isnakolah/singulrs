package main

import (
	"bankGolang/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// account, err := accounts.NewUser("Test", "Account", 20, 2000)
	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.Run()
}
