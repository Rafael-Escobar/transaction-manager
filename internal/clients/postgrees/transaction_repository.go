package postgrees

import (
	"database/sql"
	"errors"

	"github.com/transaction-manager/internal/domain"
)

type TransactionRepository struct {
	*Client
}

func NewTransactionRepository(db *Client) *TransactionRepository {
	return &TransactionRepository{db}
}

func (s *TransactionRepository) Create(transaction *domain.Transaction) (int64, error) {

	query := "INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, $4) RETURNING id;"
	result := s.db.QueryRow(query,
		transaction.AccountID,
		transaction.OperationTypeID,
		transaction.Amount,
		transaction.EventDate,
	)
	if result.Err() != nil && !errors.Is(result.Err(), sql.ErrNoRows) {
		return int64(0), result.Err()
	}
	err := result.Scan(&transaction.ID)
	if err != nil {
		return int64(0), err
	}
	return transaction.ID, nil
}
