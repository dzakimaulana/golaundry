package models

import "github.com/google/uuid"

type TransactionsItems struct {
	TransactionsID uuid.UUID `json:"transactions_id"`
	ItemsID        uuid.UUID `json:"items_id"`
	Quantity       float64   `json:"quantity"`
	Amount         int64     `json:"amount"`
	Items          *Items    `json:"items"`
}