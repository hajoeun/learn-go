package main

import (
	"fmt"

	"github.com/joeunha/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("joeun")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
