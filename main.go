package main

import (
	"bankGolang/transactions"
	"fmt"
)

func main() {
	sugar := transactions.NewItem("Sugar", map[string]float64{"Kabras": 110, "Mumias": 130}, "kg(s)")
	fmt.Printf("%T", sugar.Brands)
}
