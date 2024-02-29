package postgrees

import (
	"database/sql"
	"errors"

	"github.com/transaction-manager/internal/domain"
)

type AccountRepository struct {
	*Client
}

func NewAccountRepository(db *Client) *AccountRepository {
	return &AccountRepository{db}
}

func (s *AccountRepository) Create(account *domain.Account) (int64, error) {
	query := "INSERT INTO Accounts (document_number) VALUES ($1) RETURNING id;"
	result := s.db.QueryRow(query, account.DocumentNumber)
	if result.Err() != nil && !errors.Is(result.Err(), sql.ErrNoRows) {
		return int64(0), result.Err()
	}
	err := result.Scan(&account.ID)
	if err != nil {
		return int64(0), err
	}
	return account.ID, nil
}

func (s *AccountRepository) FindByDocumentNumber(documentNumber string) (*domain.Account, error) {
	var account domain.Account
	err := s.db.Get(&account, "SELECT id,document_number FROM Accounts WHERE document_number = $1", documentNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &account, err
}

func (s *AccountRepository) FindByID(id int64) (*domain.Account, error) {
	var account domain.Account
	err := s.db.Get(&account, "SELECT id,document_number FROM Accounts WHERE id = $1", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &account, err
}
