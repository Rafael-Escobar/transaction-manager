package domain

type Transaction struct {
	ID              string
	AccountID       int
	OperationTypeID int
	Amount          float64
	EventDate       string
}
