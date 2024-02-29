package ports

import "github.com/transaction-manager/internal/domain"

type OperationTypeRepository interface {
	FindByID(id int) (*domain.OperationType, error)
}
