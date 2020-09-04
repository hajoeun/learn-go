package nomadcoder

import (
	"errors"
	"fmt"
)

type account struct {
	owner   string
	balance int
}

var errNotEnoughMoney = errors.New("Not enough money to withdraw")

// NewAccount Function
func NewAccount(owner string) *account {
	newAccount := account{owner: owner, balance: 0}

	return &newAccount
}

// Deposit set amount
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Balance return balance of account
func (a account) Balance() int {
	return a.balance
}

// Withdraw get amount from account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNotEnoughMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner return owner
func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.owner, "'s account\nBalance: ", a.balance)
}
