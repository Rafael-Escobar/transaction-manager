package usecases

import (
	"context"

	"github.com/transaction-manager/internal/domain"
	"github.com/transaction-manager/internal/usecases/ports"
	"go.uber.org/zap"
)

type CreateTransactionUseCase struct {
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
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionRepository:   transactionRepository,
		AccountRepository:       accountRepository,
		OperationTypeRepository: operationTypeRepository,
		logger:                  logger,
	}
}

func (c *CreateTransactionUseCase) Run(ctx context.Context, transaction *domain.Transaction) (error, int64) {
	c.logger.Info("[CreateTransactionUseCase] starting", zap.Any("transaction", transaction))

	account, err := c.AccountRepository.FindByID(transaction.AccountID)
	if err != nil {
		c.logger.Error("[CreateTransactionUseCase] error finding account by ID", zap.Error(err))
		return err, 0
	}
	if account == nil {
		c.logger.Info("[CreateTransactionUseCase] account not found", zap.Any("accountID", transaction.AccountID))
		return domain.ErrIncorrectAccount, 0
	}

	operationType, err := c.OperationTypeRepository.FindByID(transaction.OperationTypeID)
	if err != nil {
		c.logger.Error("[CreateTransactionUseCase] error finding operation type by ID", zap.Error(err))
		return err, 0
	}
	if operationType == nil {
		c.logger.Info("[CreateTransactionUseCase] operation type not found", zap.Any("operationTypeID", transaction.OperationTypeID))
		return domain.ErrIncorrectOperationType, 0
	}

	transactionID, err := c.TransactionRepository.Create(transaction)
	if err != nil {
		c.logger.Error("[CreateTransactionUseCase] error creating transaction", zap.Error(err))
		return err, 0
	}
	c.logger.Info("[CreateTransactionUseCase] transaction created", zap.Any("transactionID", transactionID))
	return nil, transactionID
}
