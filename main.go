package main

import (
	"fmt"

	"github.com/joeunha/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("joeun")
	account.Deposit(10)
	res := account.Withdraw(20)
	if res != nil {
		fmt.Println(res)
	}
	fmt.Println(account)
}
