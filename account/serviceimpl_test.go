package account

import (
	"errors"
	"github.com/feryadialoi/go-error-handling-practice/errorutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceImpl_GetAccount(t *testing.T) {
	ResetAccountMap()
	repository := NewLocalRepository()
	service := NewServiceImpl(repository)

	getAccountResponse, err := service.GetAccount(GetAccountRequest{AccountNumber: "ACCOUNT-001"})
	assert.NoError(t, err)

	assert.Equal(t, "ACCOUNT-001", getAccountResponse.AccountNumber)
}

func TestServiceImpl_GetAccount_NotFoundErr(t *testing.T) {
	ResetAccountMap()
	repository := NewLocalRepository()
	service := NewServiceImpl(repository)

	_, err := service.GetAccount(GetAccountRequest{AccountNumber: "ACCOUNT-003"})
	assert.ErrorIs(t, errorutil.EntityNotFoundError, errors.Unwrap(err))
}

func TestServiceImpl_TopUp(t *testing.T) {
	ResetAccountMap()
	repository := NewLocalRepository()
	service := NewServiceImpl(repository)

	account, err := repository.FindByAccountNumber("ACCOUNT-001")
	assert.NoError(t, err)
	assert.Equal(t, float64(0), account.Balance)

	// top up 1x
	response, err := service.TopUp(TopUpRequest{
		AccountNumber: "ACCOUNT-001",
		Amount:        100_000,
	})
	assert.NoError(t, err)
	assert.Equal(t, "ACCOUNT-001", response.AccountNumber)
	assert.Equal(t, float64(100_000), response.Amount)
	assert.Equal(t, float64(100_000), response.Balance)

	account, err = repository.FindByAccountNumber("ACCOUNT-001")
	assert.NoError(t, err)
	assert.Equal(t, float64(100_000), account.Balance)

	// top up 2x
	response, err = service.TopUp(TopUpRequest{
		AccountNumber: "ACCOUNT-001",
		Amount:        100_000,
	})
	assert.NoError(t, err)
	assert.Equal(t, "ACCOUNT-001", response.AccountNumber)
	assert.Equal(t, float64(100_000), response.Amount)
	assert.Equal(t, float64(200_000), response.Balance)

	account, err = repository.FindByAccountNumber("ACCOUNT-001")
	assert.NoError(t, err)
	assert.Equal(t, float64(200_000), account.Balance)
}

func TestServiceImpl_Transfer(t *testing.T) {
	ResetAccountMap()
	repository := NewLocalRepository()
	service := NewServiceImpl(repository)

	sourceAccount, err := repository.FindByAccountNumber("ACCOUNT-001")
	assert.NoError(t, err)
	assert.Equal(t, float64(0), sourceAccount.Balance)

	destinationAccount, err := repository.FindByAccountNumber("ACCOUNT-002")
	assert.NoError(t, err)
	assert.Equal(t, float64(0), destinationAccount.Balance)

	topUpResponse, err := service.TopUp(TopUpRequest{
		AccountNumber: "ACCOUNT-001",
		Amount:        100_000,
	})
	assert.NoError(t, err)
	assert.Equal(t, "ACCOUNT-001", topUpResponse.AccountNumber)
	assert.Equal(t, float64(100_000), topUpResponse.Amount)
	assert.Equal(t, float64(100_000), topUpResponse.Balance)

	transferResponse, err := service.Transfer(TransferRequest{
		SourceAccountNumber:      "ACCOUNT-001",
		DestinationAccountNumber: "ACCOUNT-002",
		Amount:                   10_000,
	})
	assert.NoError(t, err)
	assert.Equal(t, float64(10_000), transferResponse.Amount)
	assert.Equal(t, "ACCOUNT-001", transferResponse.SourceAccountNumber)
	assert.Equal(t, "ACCOUNT-002", transferResponse.DestinationAccountNumber)

	sourceAccount, err = repository.FindByAccountNumber("ACCOUNT-001")
	assert.NoError(t, err)
	assert.Equal(t, float64(90_000), sourceAccount.Balance)

	destinationAccount, err = repository.FindByAccountNumber("ACCOUNT-002")
	assert.NoError(t, err)
	assert.Equal(t, float64(10_000), destinationAccount.Balance)
}

func TestServiceImpl_Transfer_BalanceInsufficientErr(t *testing.T) {
	ResetAccountMap()
	repository := NewLocalRepository()
	service := NewServiceImpl(repository)

	_, err := service.Transfer(TransferRequest{
		SourceAccountNumber:      "ACCOUNT-001",
		DestinationAccountNumber: "ACCOUNT-002",
		Amount:                   10_000_000,
	})
	assert.ErrorIs(t, errorutil.BalanceInsufficientError, err)
}
