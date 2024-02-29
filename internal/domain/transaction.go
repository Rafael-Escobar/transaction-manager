package domain

import (
	"errors"
	"time"
)

var (
	ErrIncorrectOperationType          = errors.New("incorrect operation type")
	ErrIncorrectAccount                = errors.New("incorrect account")
	ErrIncorrectAmountForOperationType = errors.New("incorrect amount for operation type")
)

type Transaction struct {
	ID              int64     `db:"id"`
	AccountID       int64     `db:"account_id"`
	OperationTypeID int       `db:"operation_type_id"`
	Amount          float64   `db:"amount"`
	EventDate       time.Time `db:"event_date"`
}
