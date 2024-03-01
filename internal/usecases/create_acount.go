package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases/ports"
	"go.uber.org/zap"
)

type CreateAccountUseCase interface {
	Run(ctx context.Context, account *domain.Account) (int64, error)
}
type createAccountUseCase struct {
	AccountRepository ports.AccountRepository
	logger            *zap.Logger
}

func NewCreateAccountUseCase(
	accountRepository ports.AccountRepository,
	logger *zap.Logger,
) *createAccountUseCase {
	return &createAccountUseCase{
		AccountRepository: accountRepository,
		logger:            logger,
	}
}

func (c *createAccountUseCase) Run(ctx context.Context, account *domain.Account) (int64, error) {
	c.logger.Info("[createAccountUseCase] starting", zap.Any("account", account))
	if !account.IsDocumentNumberValid() {
		c.logger.Info("[createAccountUseCase] invalid document number", zap.Any("account", account))
		return 0, domain.ErrInvalidDocumentNumber
	}
	acc, err := c.AccountRepository.FindByDocumentNumber(account.DocumentNumber)
	if err != nil {
		c.logger.Error("[createAccountUseCase] error finding account by document number", zap.Error(err))
		return 0, err
	}
	if acc != nil {
		c.logger.Info("[createAccountUseCase] account already exists", zap.Any("account", account))
		return 0, domain.ErrAccountAlreadyExists
	}
	accountID, err := c.AccountRepository.Create(account)
	if err != nil {
		c.logger.Error("[createAccountUseCase] error creating account", zap.Error(err))
		return 0, err
	}
	c.logger.Info("[createAccountUseCase] account created", zap.Any("accountID", accountID))
	return accountID, nil
}
