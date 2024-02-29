package domain

type OperationType struct {
	ID          int    `db:"id"`
	Description string `db:"description"`
}
