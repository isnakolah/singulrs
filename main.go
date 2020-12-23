package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	sugar := transactions.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	transaction1, _, _ := transactions.NewTransaction(sugar, map[string]float64{"Kabras": 110}, 1)
	fmt.Println(transaction1)
	fmt.Printf("%T", sugar.Brands)
}
