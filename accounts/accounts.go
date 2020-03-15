package accounts

// Account is struct which compose account of bank.
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
