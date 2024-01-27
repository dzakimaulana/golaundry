package models

import "github.com/google/uuid"

type Customers struct {
	ID           uuid.UUID       `gorm:"primaryKey" json:"id"`
	Name         string          `json:"name"`
	Address      string          `json:"address"`
	PhoneNumber  string          `json:"phone_number"`
	Transactions *[]Transactions `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"transactions"`
}
