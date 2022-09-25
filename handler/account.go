package handler

import (
	"github.com/feryadialoi/go-error-handling-practice/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandler struct {
	accountService account.Service
}

func NewAccountHandler(accountService account.Service) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (handler *AccountHandler) GetAccount(ctx *gin.Context) {
	accountNumber := ctx.Param("accountNumber")
	response, err := handler.accountService.GetAccount(account.GetAccountRequest{AccountNumber: accountNumber})
	if HandleIfGetAccountError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler *AccountHandler) TopUp(ctx *gin.Context) {
	var request account.TopUpRequest
	err := ctx.ShouldBindJSON(&request)
	if HandleIfShouldBindJSONError(ctx, err, &request) {
		return
	}

	response, err := handler.accountService.TopUp(request)
	if HandleIfTopUpError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler *AccountHandler) Transfer(ctx *gin.Context) {
	var request account.TransferRequest
	err := ctx.ShouldBindJSON(&request)
	if HandleIfShouldBindJSONError(ctx, err, &request) {
		return
	}

	response, err := handler.accountService.Transfer(request)
	if HandleIfTransferError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, response)
}
