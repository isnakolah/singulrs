package main

import (
	"bankGolang/accounts"
	"fmt"
)

func main() {
	daniel := accounts.New("Daniel", "Nakolah", 20, 2, -3000)
	daniel.Deposit(4000)
	daniel.Withdraw(8000)
	fmt.Println(daniel.Acc.Balance)
}
