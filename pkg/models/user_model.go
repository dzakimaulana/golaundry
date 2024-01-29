package models

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID       `gorm:"primaryKey" json:"id"`
	Username     string          `json:"username"`
	Password     string          `json:"password"`
	Role         string          `json:"role"`
	Transactions *[]Transactions `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
}

type UserRes struct {
	ID       uuid.UUID `son:"id"`
	Username string    `json:"username"`
}

type UserResByID struct {
	ID           uuid.UUID           `json:"id"`
	Username     string              `json:"username"`
	Transactions []TransactionsResUs `json:"transactions"`
}

type LoginRes struct {
	ID          uuid.UUID `json:"id"`
	AccessToken string    `json:"access_token"`
	Username    string    `json:"username"`
}
