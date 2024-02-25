package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
)

type CreateTransactionUseCase struct{}

func NewCreateTransactionUseCase() *CreateTransactionUseCase {
	return &CreateTransactionUseCase{}
}

func (*CreateTransactionUseCase) Run(ctx context.Context, transaction domain.Transaction) (error, int) {

	return nil, 0
}
