package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases/ports"
)

type CreateAccountUseCase struct {
	AccountRepository ports.AccountRepository
}

func NewCreateAccountUseCase(accountRepository ports.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountRepository: accountRepository,
	}
}

func (*CreateAccountUseCase) Run(ctx context.Context, account domain.Account) (error, int) {

	return nil, 0
}
