package models

import "github.com/google/uuid"

type Items struct {
	ID           uuid.UUID            `gorm:"primaryKey" json:"id"`
	Name         string               `json:"name"`
	Price        int64                `json:"price"`
	Unit         string               `json:"unit"`
	Duration     int64                `json:"duration"`
	Transactions *[]TransactionsItems `gorm:"foreignKey:ItemsID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

type ItemRes struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    int64     `json:"price"`
	Unit     string    `json:"unit"`
	Duration int64     `json:"duration"`
}

type ItemResByID struct {
	ID           uuid.UUID         `json:"id"`
	Name         string            `json:"name"`
	Price        int64             `json:"price"`
	Unit         string            `json:"unit"`
	Duration     int64             `json:"duration"`
	Transactions []TransItemsResIt `json:"transactions"`
}
