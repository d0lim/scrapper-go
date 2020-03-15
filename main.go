package main

import (
	"fmt"

	"github.com/imdigo/scrapper-go/accounts"
)

func main() {
	account := accounts.NewAccount("dolim")
	account.Deposit(10)
	fmt.Println(account)
}
