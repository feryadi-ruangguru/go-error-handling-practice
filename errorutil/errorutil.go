package errorutil

import "errors"

var (
	EntityNotFoundError      = errors.New("account not found")
	BalanceInsufficientError = errors.New("account balance insufficient")
)
