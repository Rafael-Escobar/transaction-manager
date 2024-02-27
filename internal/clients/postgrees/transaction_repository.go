package postgrees

import (
	"github.com/transaction-manager/internal/domain"
)

type TransactionRepository struct {
	db *Client
}

func NewTransactionRepository(db *Client) *TransactionRepository {
	return &TransactionRepository{db}
}

func (s *TransactionRepository) Create(Transaction *domain.Transaction) (int64, error) {
	result, err := s.db.db.Exec(
		"INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, $4)",
		Transaction.AccountID,
		Transaction.OperationTypeID,
		Transaction.Amount,
		Transaction.EventDate,
	)
	if err != nil {
		return result.LastInsertId()
	}
	return int64(0), err
}
