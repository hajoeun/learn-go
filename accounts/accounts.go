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
