package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	sugar := transactions.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	transaction1, _, err := transactions.NewTransaction(sugar, map[string]float64{"Kabras": 110}, 1)
	fmt.Println(transaction1, err)
	fmt.Printf("%T", sugar.Brands)
}
