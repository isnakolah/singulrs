package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	shoe := transactions.NewItem("Shoe", map[string]float64{"Balenciagas": 3000.02}, "pair(s)")
	fmt.Println(shoe.Brands)
	trans1 := transactions.NewTransaction(shoe, map[string]float64{"North Star": 300}, 3)
	trans2 := transactions.NewTransaction(shoe, map[string]float64{"North Star": 300}, 1)
	trans3 := transactions.NewTransaction(shoe, map[string]float64{"North Star": 300}, 1)

	fmt.Println(shoe.Brands)
	fmt.Println(trans1.ItemBrand, trans2.ItemBrand, trans3.ItemBrand)
}
