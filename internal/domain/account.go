package domain

import (
	"errors"

	"github.com/klassmann/cpfcnpj"
)

var (
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountAlreadyExists  = errors.New("account already exists")
	ErrInvalidDocumentNumber = errors.New("invalid document number")
)

type Account struct {
	ID             int64  `db:"id"`
	DocumentNumber string `db:"document_number"`
}

func (a *Account) IsDocumentNumberValid() bool {
	return cpfcnpj.ValidateCNPJ(a.DocumentNumber) || cpfcnpj.ValidateCPF(a.DocumentNumber)
}
