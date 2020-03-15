package main

import (
	"fmt"

	"github.com/imdigo/scrapper-go/accounts"
)

func main() {
	account := accounts.NewAccount("dolim")
	fmt.Println(account)
}
