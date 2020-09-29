package main

import (
	"fmt"

	"./accounts"
)

func main() {

	account := accounts.NewAccount("byungwook")
	account.Deposit(10)
	err := account.Withdraw(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.String())

}
