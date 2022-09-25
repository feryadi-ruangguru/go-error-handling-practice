package account

import (
	"fmt"
	"github.com/feryadialoi/go-error-handling-practice/errorutil"
)

type ServiceImpl struct {
	repository Repository
}

func NewServiceImpl(repository Repository) Service {
	return &ServiceImpl{repository: repository}
}

func (service *ServiceImpl) GetAccount(request GetAccountRequest) (response GetAccountResponse, err error) {
	account, err := service.repository.FindByAccountNumber(request.AccountNumber)
	if err != nil {
		return response, fmt.Errorf("action=service.repository.FindByAccountNumber err=%w", err)
	}

	response.ID = account.ID
	response.AccountNumber = account.AccountNumber
	response.Balance = account.Balance

	return response, err
}

func (service *ServiceImpl) TopUp(request TopUpRequest) (response TopUpResponse, err error) {
	account, err := service.repository.FindByAccountNumber(request.AccountNumber)
	if err != nil {
		return response, fmt.Errorf("action=service.repository.FindByAccountNumber accountNumber=%s err=%w",
			request.AccountNumber, err)
	}

	account.Balance = account.Balance + request.Amount

	err = service.repository.Save(account)
	if err != nil {
		return response, fmt.Errorf("action=service.repository.Save err=%w", err)
	}

	response.AccountNumber = request.AccountNumber
	response.Amount = request.Amount
	response.Balance = account.Balance

	return response, err
}

func (service *ServiceImpl) Transfer(request TransferRequest) (response TransferResponse, err error) {
	sourceAccount, err := service.repository.FindByAccountNumber(request.SourceAccountNumber)
	if err != nil {
		return response, fmt.Errorf("action=service.repository.FindByAccountNumber sourceAccountNumber=%s err=%w",
			request.SourceAccountNumber, err)
	}

	if sourceAccount.Balance < request.Amount {
		return response, errorutil.BalanceInsufficientError
	}

	destinationAccount, err := service.repository.FindByAccountNumber(request.DestinationAccountNumber)
	if err != nil {
		return response, fmt.Errorf("action=service.repository.FindByAccountNumber destinationAccountNumber=%s err=%w",
			request.DestinationAccountNumber, err)
	}

	sourceAccount.Balance = sourceAccount.Balance - request.Amount
	destinationAccount.Balance = destinationAccount.Balance + request.Amount

	err = service.repository.Save(sourceAccount)
	if err != nil {
		return response, fmt.Errorf("action=service.repository.Save sourceAccount err=%w", err)
	}

	err = service.repository.Save(destinationAccount)
	if err != nil {
		return response, fmt.Errorf("action=service.repository.Save destinationAccount err=%w", err)
	}

	response.SourceAccountNumber = request.SourceAccountNumber
	response.DestinationAccountNumber = request.DestinationAccountNumber
	response.Amount = request.Amount

	return response, err
}
