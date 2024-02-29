package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases"
	"go.uber.org/zap"
)

type Account struct {
	createAccount *usecases.CreateAccountUseCase
	getAccount    *usecases.GetAccountUseCase
	logger        *zap.Logger
}

func NewAccountHandler(
	createAccount *usecases.CreateAccountUseCase,
	getAccount *usecases.GetAccountUseCase,
	logger *zap.Logger,
) *Account {
	return &Account{
		createAccount: createAccount,
		getAccount:    getAccount,
		logger:        logger,
	}
}

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

// CreateAccount
// @Summary Create an account
// @Description	Endpoint for creating an account
// @Tags github.com/rafael-escobar/transaction-manager/
// @Produce json
// @Success 200
// @Failure	400	{object}	map[string]string
// @Failure	500	{object}	map[string]string
// @Router /v1/accounts [post]
func (a *Account) CreateAccountHandler(ctx *gin.Context) {
	a.logger.Info("[CreateAccountHandler] starting")
	defer a.logger.Info("[CreateAccountHandler] ending")

	requestBody := CreateAccountRequest{}
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	account := a.mapCreateAccountRequest(requestBody)
	err, accountID := a.createAccount.Run(ctx, account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error creating account")
		return
	}
	ctx.JSON(http.StatusOK, a.mapCreateAccountResponse(accountID))
}

// GetAccount
// @Summary Get an account
// @Description	Endpoint for getting an account
// @Tags github.com/rafael-escobar/transaction-manager/
// @Produce json
// @Success 200
// @Failure	400	{object}	map[string]string
// @Failure	500	{object}	map[string]string
// @Router /v1/accounts/{id} [Get]
func (a *Account) GetAccountHandler(ctx *gin.Context) {
	a.logger.Info("[GetAccountHandler] starting")
	defer a.logger.Info("[GetAccountHandler] ending")
	pathParam := ctx.Param("id")
	accountID, err := strconv.ParseInt(pathParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid account id")
		return
	}
	err, account := a.getAccount.Run(ctx, accountID)
	if errors.Is(err, domain.ErrAccountNotFound) {
		ctx.JSON(http.StatusNotFound, "Account not found")
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error getting account")
		return
	}
	ctx.JSON(http.StatusOK, a.mapGetAccountResponse(account))
}

type CreateAccountResponse struct {
	AccountID int64 `json:"account_id"`
}
type GetAccountResponse struct {
	AccountID      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (a *Account) mapCreateAccountResponse(accountID int64) CreateAccountResponse {
	return CreateAccountResponse{
		AccountID: accountID,
	}
}

func (a *Account) mapGetAccountResponse(account *domain.Account) GetAccountResponse {
	return GetAccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}

func (a *Account) mapCreateAccountRequest(requestBody CreateAccountRequest) *domain.Account {
	return &domain.Account{
		DocumentNumber: requestBody.DocumentNumber,
	}
}
