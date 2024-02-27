package postgrees

import (
	"github.com/transaction-manager/internal/domain"
)

type AccountRepository struct {
	*Client
}

func NewAccountRepository(db *Client) *AccountRepository {
	return &AccountRepository{db}
}

func (s *AccountRepository) Create(Account *domain.Account) (int64, error) {
	result, err := s.db.Exec("INSERT INTO Accounts (document_number) VALUES ($1)", Account.DocumentNumber)
	if err != nil {
		return result.LastInsertId()
	}
	return int64(0), err
}

func (s *AccountRepository) FindByDocumentNumber(documentNumber string) (*domain.Account, error) {
	var account domain.Account
	err := s.db.Get(&account, "SELECT * FROM Accounts WHERE document_number = $1", documentNumber)
	return &account, err
}

func (s *AccountRepository) FindByID(id int64) (*domain.Account, error) {
	var account domain.Account
	err := s.db.Get(&account, "SELECT * FROM Accounts WHERE id = $1", id)
	return &account, err
}
