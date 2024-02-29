package domain

import "errors"

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrAccountAlreadyExists = errors.New("account already exists")
)

type Account struct {
	ID             int64  `db:"id"`
	DocumentNumber string `db:"document_number"`
}
