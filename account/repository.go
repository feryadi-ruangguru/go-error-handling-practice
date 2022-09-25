package account

type Repository interface {
	FindByID(id int64) (*Account, error)
	FindByAccountNumber(accountNumber string) (*Account, error)
	Save(account *Account) error
}
