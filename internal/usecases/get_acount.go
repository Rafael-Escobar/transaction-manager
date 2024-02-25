package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
)

type GetAccountUseCase struct{}

func NewGetAccountUseCase() *GetAccountUseCase {
	return &GetAccountUseCase{}
}

func (*GetAccountUseCase) Run(ctx context.Context, accountID int) (error, domain.Account) {

	return nil, domain.Account{}
}
