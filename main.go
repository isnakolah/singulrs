package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	shoes := transactions.NewItem("Snickers", map[string]int{"Balenciagas": 5000})
	fmt.Println(shoes)
}
