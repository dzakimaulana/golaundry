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

type TransactionRepository interface {
	AddTs(ctx context.Context, t *models.Transactions) (*models.Transactions, error)
	GetAllTs(ctx context.Context) (*[]models.Transactions, error)
	GetTsByID(ctx context.Context, id string) (*models.Transactions, error)
	FindCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error)
	GetItemByID(ctx context.Context, id string) (*models.Items, error)
	AddTransItems(ctx context.Context, ti *[]models.TransactionsItems) (*[]models.TransactionsItems, error)
}

type TransactionService interface {
	AddTs(ctx context.Context, t *TransactionReq) (*models.TransactionsRes, error)
	GetAllTs(ctx context.Context) (*[]models.TransactionsRes, error)
	GetTsById(ctx context.Context, id string) (*models.TransactionResByID, error)
}
