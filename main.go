package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	sugar, itemErr := transactions.NewItem("", map[string]float64{"Kabras": 110, "Mumias": 130}, "")
	// transaction1, _, err := transactions.NewTransaction(sugar, map[string]float64{"Kabras": 110}, 1)
	// fmt.Println(transaction1, err)
	// fmt.Printf("%T", sugar.Brands)

	// daniel, err := accounts.NewUser("Daniel", "Nakolah", 2, 1000)

	// fmt.Println(daniel, err)
	fmt.Println(sugar, itemErr)
}
