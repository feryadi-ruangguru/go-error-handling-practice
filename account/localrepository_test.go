package account

import (
	"github.com/feryadialoi/go-error-handling-practice/errorutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalRepository_FindByAccountNumber(t *testing.T) {
	ResetAccountMap()
	repository := NewLocalRepository()

	account, err := repository.FindByAccountNumber("ACCOUNT-001")
	assert.NoError(t, err)

	assert.NotNil(t, account)
	assert.Equal(t, "ACCOUNT-001", account.AccountNumber)
}

func TestLocalRepository_FindByAccountNumber_NotFoundErr(t *testing.T) {
	ResetAccountMap()
	repository := NewLocalRepository()

	account, err := repository.FindByAccountNumber("ACCOUNT-003")
	assert.ErrorIs(t, errorutil.EntityNotFoundError, err)
	assert.Nil(t, account)
}
