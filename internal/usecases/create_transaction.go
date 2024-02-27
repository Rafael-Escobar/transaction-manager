package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases/ports"
)

type CreateTransactionUseCase struct {
	TransactionRepository ports.TransactionRepository
	AccountRepository     ports.AccountRepository
}

func NewCreateTransactionUseCase(
	transactionRepository ports.TransactionRepository,
	accountRepository ports.AccountRepository,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionRepository: transactionRepository,
		AccountRepository:     accountRepository,
	}
}

func (*CreateTransactionUseCase) Run(ctx context.Context, transaction domain.Transaction) (error, int) {

	return nil, 0
}
