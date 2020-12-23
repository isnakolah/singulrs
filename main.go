package main

import (
	"bankGolang/accounts"
	"fmt"
)

func main() {
	daniel, _ := accounts.NewUser("Daniel", "Nakolah", 20, 3000)

	withdraw := daniel.Withdraw
	if withdraw(30000) != nil {
		fmt.Println(withdraw(30000), daniel.Acc.Balance)
	}

}
