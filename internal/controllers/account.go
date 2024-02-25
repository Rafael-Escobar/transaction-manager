package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases"
)

type Account struct {
	createAccount *usecases.CreateAccountUseCase
	getAccount    *usecases.GetAccountUseCase
}

func NewAccountHandler(
	createAccount *usecases.CreateAccountUseCase,
	getAccount *usecases.GetAccountUseCase,
) *Account {
	return &Account{
		createAccount: createAccount,
		getAccount:    getAccount,
	}
}

func (a *Account) CreateAccountHandler(ctx *gin.Context) {
	account := domain.Account{}
	err, accountID := a.createAccount.Run(ctx, account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error creating account")
		return
	}
	ctx.JSON(http.StatusOK, a.mapCreateAccount(accountID))
}

func (a *Account) GetAccountHandler(ctx *gin.Context) {
	accountID := 0
	err, account := a.getAccount.Run(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error getting account")

		return
	}
	ctx.JSON(http.StatusOK, a.mapGetAccount(account))
}

type CreateAccountResponse struct {
	AccountID int `json:"account_id"`
}
type GetAccountResponse struct {
	AccountID      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (a *Account) mapCreateAccount(accountID int) CreateAccountResponse {
	return CreateAccountResponse{
		AccountID: accountID,
	}
}

func (a *Account) mapGetAccount(account domain.Account) GetAccountResponse {
	return GetAccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}
