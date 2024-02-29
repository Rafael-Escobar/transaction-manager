package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases"
	"go.uber.org/zap"
)

type Transaction struct {
	createTransaction *usecases.CreateTransactionUseCase
	logger            *zap.Logger
}

func NewTransactionHandler(
	createTransaction *usecases.CreateTransactionUseCase,
	logger *zap.Logger,
) *Transaction {
	return &Transaction{
		createTransaction: createTransaction,
		logger:            logger,
	}
}

type CreateTransactionRequest struct {
	AccountID       int64   `json:"account_id"`
	Amount          float64 `json:"amount"`
	OperationTypeID int     `json:"operation_type_id"`
}

// CreateTransaction
// @Summary Create an transaction
// @Description	Endpoint for creating an transaction
// @Tags github.com/rafael-escobar/transaction-manager/
// @Produce json
// @Success 200
// @Failure	400	{object}	map[string]string
// @Failure	500	{object}	map[string]string
// @Router /v1/transactions [post]
func (t *Transaction) CreateTransactionHandler(ctx *gin.Context) {
	t.logger.Info("[CreateTransactionHandler] starting")
	defer t.logger.Info("[CreateTransactionHandler] ending")

	var body CreateTransactionRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	transaction := t.mapCreateTransactionRequest(body)
	err, transactionID := t.createTransaction.Run(ctx, transaction)
	if errors.Is(err, domain.ErrIncorrectAccount) {
		ctx.JSON(http.StatusBadRequest, "Incorrect account")
		return
	}
	if errors.Is(err, domain.ErrIncorrectOperationType) {
		ctx.JSON(http.StatusBadRequest, "Incorrect operation type")
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error creating transaction")
		return
	}
	ctx.JSON(http.StatusOK, t.mapCreateTransactionResponse(transactionID))
}

type CreateTransactionResponse struct {
	TransactionID int64 `json:"transaction_id"`
}

func (t *Transaction) mapCreateTransactionResponse(transactionID int64) CreateTransactionResponse {
	return CreateTransactionResponse{
		TransactionID: transactionID,
	}
}

func (t *Transaction) mapCreateTransactionRequest(body CreateTransactionRequest) *domain.Transaction {
	return &domain.Transaction{
		AccountID:       body.AccountID,
		Amount:          body.Amount,
		OperationTypeID: body.OperationTypeID,
	}
}
