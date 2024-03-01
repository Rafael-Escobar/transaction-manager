package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases/ports"
	"go.uber.org/zap"
)

type GetAccountUseCase interface {
	Run(ctx context.Context, accountID int64) (*domain.Account, error)
}
type getAccountUseCase struct {
	AccountRepository ports.AccountRepository
	logger            *zap.Logger
}

func NewGetAccountUseCase(
	accountRepository ports.AccountRepository,
	logger *zap.Logger,
) *getAccountUseCase {
	return &getAccountUseCase{
		AccountRepository: accountRepository,
		logger:            logger,
	}
}

func (g *getAccountUseCase) Run(ctx context.Context, accountID int64) (*domain.Account, error) {
	g.logger.Info("[getAccountUseCase] starting", zap.Any("accountID", accountID))
	account, err := g.AccountRepository.FindByID(accountID)
	if err != nil {
		g.logger.Error("[getAccountUseCase] error finding account by ID", zap.Error(err))
		return nil, err
	}
	if account == nil {
		g.logger.Info("[getAccountUseCase] account not found", zap.Any("accountID", accountID))
		return nil, domain.ErrAccountNotFound
	}
	g.logger.Info("[getAccountUseCase] account found", zap.Any("account", account))
	return account, nil
}
