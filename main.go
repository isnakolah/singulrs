package main

import (
	"bankGolang/accounts"
	"fmt"
)

func main() {
	daniel, err := accounts.New("Daniel", "Nakolah", 20, 2, 100)

	fmt.Println(err)
	fmt.Println(daniel.Acc.Balance)
}
