package transitems

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
)

type TransItemsReq struct {
	ItemsID  uuid.UUID `json:"items_id"`
	Quantity float64   `json:"quantity"`
}

type TransItemsRes struct {
	Quantity float64       `json:"quantity"`
	Amount   int64         `json:"amount"`
	Items    *models.Items `json:"items"`
}

type TransItemsRepository interface {
	AddTransItems(ctx context.Context, ti *[]models.TransactionsItems) (*[]models.TransactionsItems, error)
}
