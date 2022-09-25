package account

type GetAccountRequest struct {
	AccountNumber string `json:"accountNumber" binding:"required"`
}

type GetAccountResponse struct {
	ID            int64   `json:"id"`
	AccountNumber string  `json:"accountNumber"`
	Balance       float64 `json:"balance"`
}

type TopUpRequest struct {
	AccountNumber string  `json:"accountNumber" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}

type TopUpResponse struct {
	AccountNumber string  `json:"accountNumber"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
}

type TransferRequest struct {
	SourceAccountNumber      string  `name:"sourceAccountNumber" json:"sourceAccountNumber" binding:"required"`
	DestinationAccountNumber string  `json:"destinationAccountNumber" binding:"required"`
	Amount                   float64 `json:"amount" binding:"required,min=10"`
}

type TransferResponse struct {
	SourceAccountNumber      string  `json:"sourceAccountNumber"`
	DestinationAccountNumber string  `json:"destinationAccountNumber"`
	Amount                   float64 `json:"amount"`
}
