package main

import (
	"bankGolang/accounts"
	"fmt"
)

func main() {
	daniel, err := accounts.New("Daniel", "Nakolah", 20, 100)

	fmt.Println(err)
	fmt.Println(daniel.Acc.Balance)
	fmt.Println(daniel.Acc.Number)
}
