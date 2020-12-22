package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	shoe := transactions.NewItem("Shoe", map[string]float64{"Balenciagas": 3000.02}, "pair(s)")
	trans1 := transactions.NewTransaction(shoe, map[string]float64{"North Star": 300}, 3)

	fmt.Println(trans1.ItemBrand)
}
