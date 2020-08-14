package accounts

import "errors"

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
