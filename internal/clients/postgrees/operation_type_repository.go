package postgrees

import (
	"database/sql"
	"errors"

	"github.com/transaction-manager/internal/domain"
)

// OperationTypeRepository is a repository for the operation type entity
type OperationTypeRepository struct {
	*Client
}

// NewOperationTypeRepository creates a new instance of OperationTypeRepository
func NewOperationTypeRepository(db *Client) *OperationTypeRepository {
	return &OperationTypeRepository{db}
}

// FindByID finds an operation type by its id
func (s *OperationTypeRepository) FindByID(id int) (*domain.OperationType, error) {
	operationType := &domain.OperationType{}
	err := s.db.Get(operationType, "SELECT id, description FROM operation_types WHERE id = $1", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return operationType, nil
}
