package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner   string
	balance int
}

var errNomoney = errors.New("cant withdraw you are poor zz")

//NewAccount creates account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount balance
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from your balance
func (a *Account) Withdraw(amount int) error {
	if amount > a.balance {
		return errors.New("can`t withdraw you are poor -zz")
	}
	a.balance -= amount
	return nil
}

//ChangeOwner change owner`s name
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//String show about account
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas ", a.Balance())
}

//Owner show owner
func (a Account) Owner() string {
	return a.owner
}

//Balance show your balance
func (a Account) Balance() int {
	return a.balance
}
