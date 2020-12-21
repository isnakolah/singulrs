package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	shoes := transactions.NewItem("Snickers", map[string]float64{"Balenciagas": 5000, "North Star": 1000})
	fmt.Println(shoes)
}
