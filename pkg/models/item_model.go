package models

import "github.com/google/uuid"

type Items struct {
	ID           uuid.UUID            `gorm:"primaryKey" json:"id"`
	Name         string               `json:"name"`
	Price        int64                `json:"price"`
	Unit         string               `json:"unit"`
	Duration     int64                `json:"duration"`
	Transactions *[]TransactionsItems `gorm:"foreignKey:ItemsID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"transactions"`
}
