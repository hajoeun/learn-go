package main

import (
	"fmt"

	"github.com/joeunha/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("joeun")
	fmt.Println(account)
}
