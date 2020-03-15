package main

import (
	"fmt"

	"github.com/imdigo/scrapper-go/accounts"
)

func main() {
	account := accounts.NewAccount("dolim")
	account.Deposit(10)
	fmt.Println(account.Balance())
	if err := account.Withdraw(20); err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
