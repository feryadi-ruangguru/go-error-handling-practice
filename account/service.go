package account

type Service interface {
	GetAccount(request GetAccountRequest) (GetAccountResponse, error)
	TopUp(request TopUpRequest) (TopUpResponse, error)
	Transfer(request TransferRequest) (TransferResponse, error)
}
