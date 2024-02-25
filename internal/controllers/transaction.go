package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases"
)

type Transaction struct {
	createTransaction *usecases.CreateTransactionUseCase
}

func NewTransactionHandler(
	createTransaction *usecases.CreateTransactionUseCase,
) *Transaction {
	return &Transaction{
		createTransaction: createTransaction,
	}
}

func (t *Transaction) CreateTransactionHandler(ctx *gin.Context) {
	transaction := domain.Transaction{}
	err, transactionID := t.createTransaction.Run(ctx, transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error creating transaction")
		return
	}
	ctx.JSON(http.StatusOK, t.mapCreateTransaction(transactionID))
}

type CreateTransactionResponse struct {
	TransactionID int `json:"transaction_id"`
}

func (t *Transaction) mapCreateTransaction(transactionID int) CreateTransactionResponse {
	return CreateTransactionResponse{
		TransactionID: transactionID,
	}
}
