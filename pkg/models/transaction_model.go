package models

import "github.com/google/uuid"

type Transactions struct {
	ID         uuid.UUID            `gorm:"primaryKey" json:"id"`
	CustomerID uuid.UUID            `json:"customer_id"`
	UserID     uuid.UUID            `json:"user_id"`
	TimeIn     int64                `json:"time_in"`
	TimeOut    int64                `json:"time_out"`
	Total      int64                `json:"total"`
	Customer   *Customers           `json:"customer"`
	User       *User                `json:"user"`
	Items      *[]TransactionsItems `gorm:"foreignKey:TransactionsID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"transactions_items"`
}
