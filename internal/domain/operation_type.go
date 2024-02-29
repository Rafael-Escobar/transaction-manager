package domain

type OperationType struct {
	ID          int    `db:"id"`
	Description string `db:"description"`
	IsDebit     bool   `db:"is_debit"`
}

func (o *OperationType) IsAmountValid(amount float64) bool {
	zero := 0.0
	if amount == zero {
		return false
	}
	if o.IsDebit && amount > zero {
		return false
	}
	if !o.IsDebit && amount < zero {
		return false
	}
	return true
}
