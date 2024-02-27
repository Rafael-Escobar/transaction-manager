package ports

import "github.com/transaction-manager/internal/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) (int64, error)
}
