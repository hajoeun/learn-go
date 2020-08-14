package accounts

type account struct {
	owner   string
	balance int
}

// NewAccount Function
func NewAccount(owner string) *account {
	newAccount := account{owner: owner, balance: 0}

	return &newAccount
}

// Deposit set amount
func (a account) Deposit(amount int) {
	a.balance += amount
}

// Balance return balance of account
func (a account) Balance() int {
	return a.balance
}
