package domain

type Transaction struct {
	ID             string
	AccountID      int
	OperationTypID int
	Amount         float64
	EventDate      string
}
