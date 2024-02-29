package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases/ports"
	"go.uber.org/zap"
)

type CreateAccountUseCase struct {
	AccountRepository ports.AccountRepository
	logger            *zap.Logger
}

func NewCreateAccountUseCase(
	accountRepository ports.AccountRepository,
	logger *zap.Logger,
) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountRepository: accountRepository,
		logger:            logger,
	}
}

func (c *CreateAccountUseCase) Run(ctx context.Context, account *domain.Account) (int64, error) {
	c.logger.Info("[CreateAccountUseCase] starting", zap.Any("account", account))
	acc, err := c.AccountRepository.FindByDocumentNumber(account.DocumentNumber)
	if err != nil {
		c.logger.Error("[CreateAccountUseCase] error finding account by document number", zap.Error(err))
		return 0, err
	}
	if acc != nil {
		c.logger.Info("[CreateAccountUseCase] account already exists", zap.Any("account", account))
		return 0, domain.ErrAccountAlreadyExists
	}
	if !account.IsDocumentNumberValid() {
		c.logger.Info("[CreateAccountUseCase] invalid document number", zap.Any("account", account))
		return 0, domain.ErrInvalidDocumentNumber
	}
	accountID, err := c.AccountRepository.Create(account)
	if err != nil {
		c.logger.Error("[CreateAccountUseCase] error creating account", zap.Error(err))
		return 0, err
	}
	c.logger.Info("[CreateAccountUseCase] account created", zap.Any("accountID", accountID))
	return accountID, nil
}
