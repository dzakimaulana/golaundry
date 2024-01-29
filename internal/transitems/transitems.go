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

type TransItemsRepository interface {
	AddTransItems(ctx context.Context, ti *[]models.TransactionsItems) (*[]models.TransactionsItems, error)
}
