package items

import (
	"context"

	"github.com/google/uuid"
)

type Items struct {
	ID    uuid.UUID `gorm:"primaryKey" json:"id"`
	Name  string    `json:"name"`
	Price int64     `json:"price"`
	Unit  string    `json:"unit"`
	Time  int64     `json:"time"`
}

type AddReq struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Unit  string `json:"unit"`
	Time  int64  `json:"time"`
}

type AddRes struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price int64     `json:"price"`
	Unit  string    `json:"unit"`
	Time  int64     `json:"time"`
}

type GetRes struct {
	Items *[]Items `json:"items"`
}

type UpdateReq struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price int64     `json:"price"`
	Unit  string    `json:"unit"`
	Time  int64     `json:"time"`
}

type UpdateRes struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price int64     `json:"price"`
	Unit  string    `json:"unit"`
	Time  int64     `json:"time"`
}

type DelReq struct {
	ID uuid.UUID `json:"id"`
}

type ItemRepository interface {
	AddItem(ctx context.Context, item *Items) (*Items, error)
	GetItem(ctx context.Context, name string) (*[]Items, error)
	UpdateItem(ctx context.Context, item *Items) (*Items, error)
	DeleteItem(ctx context.Context, id string) error
}

type ItemService interface {
	AddItem(ctx context.Context, req *AddReq) (*AddRes, error)
	GetItem(ctx context.Context, name string) (*GetRes, error)
	UpdateItem(ctx context.Context, req *UpdateReq) (*UpdateRes, error)
	DeleteItem(ctx context.Context, req *DelReq) error
}
