package main

import (
	"github.com/feryadialoi/go-error-handling-practice/account"
	"github.com/feryadialoi/go-error-handling-practice/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	accountRepository := account.NewLocalRepository()
	accountService := account.NewServiceImpl(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	engine.Group("/accounts").
		GET("/:accountNumber", accountHandler.GetAccount).
		POST("/top-up", accountHandler.TopUp).
		POST("/transfer", accountHandler.Transfer)

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
