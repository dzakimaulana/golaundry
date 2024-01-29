package items

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
)

type ItemReq struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Unit     string `json:"unit"`
	Duration int64  `json:"duration"`
}

type UpdateReq struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    int64     `json:"price"`
	Unit     string    `json:"unit"`
	Duration int64     `json:"duration"`
}

type DelReq struct {
	ID uuid.UUID `json:"id"`
}

type ItemRepository interface {
	AddItem(ctx context.Context, item *models.Items) (*models.Items, error)
	GetItem(ctx context.Context, name string) (*[]models.Items, error)
	GetItemByID(ctx context.Context, id string) (*models.Items, error)
	UpdateItem(ctx context.Context, item *models.Items) (*models.Items, error)
	DeleteItem(ctx context.Context, id string) error
}

type ItemService interface {
	AddItem(ctx context.Context, req *ItemReq) (*models.ItemRes, error)
	GetItem(ctx context.Context, name string) (*[]models.ItemRes, error)
	GetItemByID(ctx context.Context, id string) (*models.ItemResByID, error)
	UpdateItem(ctx context.Context, req *UpdateReq) (*models.ItemRes, error)
	DeleteItem(ctx context.Context, req *DelReq) error
}
