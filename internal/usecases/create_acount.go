package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
)

type CreateAccountUseCase struct{}

func NewCreateAccountUseCase() *CreateAccountUseCase {
	return &CreateAccountUseCase{}
}

func (*CreateAccountUseCase) Run(ctx context.Context, account domain.Account) (error, int) {

	return nil, 0
}
