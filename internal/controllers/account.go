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
// @Param requestBody body CreateAccountRequest true "Request body"
// @Produce json
// @Success 200 {object} CreateAccountResponse
// @Failure	400	{object}	ResponseError
// @Failure	500	{object}	ResponseError
// @Router /v1/accounts [post]
func (a *Account) CreateAccountHandler(ctx *gin.Context) {
	a.logger.Info("[CreateAccountHandler] starting")
	defer a.logger.Info("[CreateAccountHandler] ending")

	requestBody := CreateAccountRequest{}
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, a.mapResponseError("Invalid request body"))
		return
	}
	account := a.mapCreateAccountRequest(requestBody)
	accountID, err := a.createAccount.Run(ctx, account)
	if errors.Is(err, domain.ErrAccountAlreadyExists) {
		ctx.JSON(http.StatusConflict, a.mapResponseError("Account already exists"))
		return
	}
	if errors.Is(err, domain.ErrInvalidDocumentNumber) {
		ctx.JSON(http.StatusBadRequest, a.mapResponseError("Invalid document number"))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, a.mapResponseError("Error creating account"))
		return
	}
	ctx.JSON(http.StatusOK, a.mapCreateAccountResponse(accountID))
}

// GetAccount
// @Summary Get an account
// @Description	Endpoint for getting an account
// @Tags github.com/rafael-escobar/transaction-manager/
// @Param id path int true "Account ID"
// @Produce json
// @Success 200	{object}	GetAccountResponse
// @Failure	400	{object}	ResponseError
// @Failure	500	{object}	ResponseError
// @Router /v1/accounts/{id} [Get]
func (a *Account) GetAccountHandler(ctx *gin.Context) {
	a.logger.Info("[GetAccountHandler] starting")
	defer a.logger.Info("[GetAccountHandler] ending")
	pathParam := ctx.Param("id")
	accountID, err := strconv.ParseInt(pathParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, a.mapResponseError("Invalid account id"))
		return
	}
	account, err := a.getAccount.Run(ctx, accountID)
	if errors.Is(err, domain.ErrAccountNotFound) {
		ctx.JSON(http.StatusNotFound, a.mapResponseError("Account not found"))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, a.mapResponseError("Error getting account"))
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

type ResponseError struct {
	Message string `json:"message"`
}

func (a *Account) mapResponseError(messageError string) ResponseError {
	return ResponseError{
		Message: messageError,
	}
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
