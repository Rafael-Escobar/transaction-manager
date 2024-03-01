package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases/ports"
	"go.uber.org/zap"
)

type CreateTransactionUseCase interface {
	Run(ctx context.Context, transaction *domain.Transaction) (int64, error)
}

type createTransactionUseCase struct {
	TransactionRepository   ports.TransactionRepository
	AccountRepository       ports.AccountRepository
	OperationTypeRepository ports.OperationTypeRepository
	logger                  *zap.Logger
}

func NewCreateTransactionUseCase(
	transactionRepository ports.TransactionRepository,
	accountRepository ports.AccountRepository,
	operationTypeRepository ports.OperationTypeRepository,
	logger *zap.Logger,
) *createTransactionUseCase {
	return &createTransactionUseCase{
		TransactionRepository:   transactionRepository,
		AccountRepository:       accountRepository,
		OperationTypeRepository: operationTypeRepository,
		logger:                  logger,
	}
}

func (c *createTransactionUseCase) Run(ctx context.Context, transaction *domain.Transaction) (int64, error) {
	c.logger.Info("[createTransactionUseCase] starting", zap.Any("transaction", transaction))

	account, err := c.AccountRepository.FindByID(transaction.AccountID)
	if err != nil {
		c.logger.Error("[createTransactionUseCase] error finding account by ID", zap.Error(err))
		return 0, err
	}
	if account == nil {
		c.logger.Info("[createTransactionUseCase] account not found", zap.Any("accountID", transaction.AccountID))
		return 0, domain.ErrInvalidAccount
	}

	operationType, err := c.OperationTypeRepository.FindByID(transaction.OperationTypeID)
	if err != nil {
		c.logger.Error("[createTransactionUseCase] error finding operation type by ID", zap.Error(err))
		return 0, err
	}
	if operationType == nil {
		c.logger.Info("[createTransactionUseCase] operation type not found", zap.Any("operationTypeID", transaction.OperationTypeID))
		return 0, domain.ErrInvalidOperationType
	}
	c.logger.Info("[createTransactionUseCase] operation type found", zap.Any("operationType", operationType))

	if !operationType.IsAmountValid(transaction.Amount) {
		c.logger.Info("[createTransactionUseCase] invalid amount for operation type",
			zap.Any("amount", transaction.Amount),
			zap.Any("operationType", operationType),
		)
		return 0, domain.ErrInvalidAmountForOperationType
	}

	transactionID, err := c.TransactionRepository.Create(transaction)
	if err != nil {
		c.logger.Error("[createTransactionUseCase] error creating transaction", zap.Error(err))
		return 0, err
	}
	c.logger.Info("[createTransactionUseCase] transaction created", zap.Any("transactionID", transactionID))
	return transactionID, nil
}
