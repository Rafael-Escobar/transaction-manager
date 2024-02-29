package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidOperationType          = errors.New("invalid operation type")
	ErrInvalidAccount                = errors.New("invalid account")
	ErrInvalidAmountForOperationType = errors.New("invalid amount for operation type")
)

type Transaction struct {
	ID              int64     `db:"id"`
	AccountID       int64     `db:"account_id"`
	OperationTypeID int       `db:"operation_type_id"`
	Amount          float64   `db:"amount"`
	EventDate       time.Time `db:"event_date"`
}
