package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID       `gorm:"primaryKey" json:"id"`
	Username     string          `json:"username"`
	Password     string          `json:"password"`
	Transactions *[]Transactions `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"transactions"`
}
