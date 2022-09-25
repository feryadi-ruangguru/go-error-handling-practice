package account

import "github.com/feryadialoi/go-error-handling-practice/errorutil"

var accountMap map[int64]*Account

func init() {
	accountMap = map[int64]*Account{
		1: {ID: 1, AccountNumber: "ACCOUNT-001", Balance: 0},
		2: {ID: 2, AccountNumber: "ACCOUNT-002", Balance: 0},
	}
}

func ResetAccountMap() {
	accountMap = map[int64]*Account{
		1: {ID: 1, AccountNumber: "ACCOUNT-001", Balance: 0},
		2: {ID: 2, AccountNumber: "ACCOUNT-002", Balance: 0},
	}
}

type LocalRepository struct {
}

func NewLocalRepository() Repository {
	return &LocalRepository{}
}

func (repository *LocalRepository) FindByID(id int64) (*Account, error) {
	if account, ok := accountMap[id]; ok {
		return account, nil
	}
	return nil, errorutil.EntityNotFoundError
}

func (repository *LocalRepository) FindByAccountNumber(accountNumber string) (*Account, error) {
	for _, account := range accountMap {
		if accountNumber == account.AccountNumber {
			return account, nil
		}
	}
	return nil, errorutil.EntityNotFoundError
}

func (repository *LocalRepository) Save(account *Account) error {
	if account.ID < 1 {
		lastID := repository.generateID()
		account.ID = lastID
		accountMap[lastID] = account
		return nil
	}
	accountMap[account.ID] = account
	return nil
}

// private method

func (repository *LocalRepository) generateID() int64 {
	var lastID = int64(len(accountMap) + 1)
	for {
		if _, ok := accountMap[lastID]; !ok {
			break
		}
		lastID++
	}
	return lastID
}
