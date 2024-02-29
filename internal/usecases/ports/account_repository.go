package ports

import "github.com/transaction-manager/internal/domain"

type AccountRepository interface {
	FindByID(id int64) (*domain.Account, error)
	FindByDocumentNumber(documentNumber string) (*domain.Account, error)
	Create(account *domain.Account) (int64, error)
}
