package transactions

import (
	"context"

	"github.com/dzakimaulana/golaundry/internal/transitems"
	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
)

type TransactionReq struct {
	UserID   uuid.UUID                  `json:"user_id"`
	Customer models.Customers           `json:"customer"`
	Items    []transitems.TransItemsReq `json:"items"`
}

type TransactionRes struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	UserID     uuid.UUID `json:"user_id"`
	TimeIn     int64     `json:"time_in"`
	TimeOut    int64     `json:"time_out"`
	Total      int64     `json:"total"`
}

type TransactionByID struct {
	ID       uuid.UUID                  `gorm:"primaryKey" json:"id"`
	Customer models.Customers           `gorm:"foreignKey:ID" json:"customers"`
	Items    []models.TransactionsItems `gorm:"foreignKey:ID" json:"items"`
	TimeIn   int64                      `json:"time_in"`
	TimeOut  int64                      `json:"time_out"`
	Total    int64                      `json:"total"`
}

type GetAllRes struct {
	Transactions *[]TransactionRes `json:"transactions"`
}

type TransactionRepository interface {
	AddTs(ctx context.Context, t *models.Transactions) (*models.Transactions, error)
	GetAllTs(ctx context.Context) (*[]models.Transactions, error)
	GetTsByID(ctx context.Context, id string) (*models.Transactions, error)
	FindCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error)
	GetItemByID(ctx context.Context, id string) (*models.Items, error)
	AddTransItems(ctx context.Context, ti *[]models.TransactionsItems) (*[]models.TransactionsItems, error)
}

type TransactionService interface {
	AddTs(ctx context.Context, t *TransactionReq) (*TransactionRes, error)
	GetAllTs(ctx context.Context) (*GetAllRes, error)
	GetTsById(ctx context.Context, id string) (*TransactionByID, error)
}
