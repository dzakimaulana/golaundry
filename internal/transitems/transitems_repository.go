package transitems

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) TransItemsRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) AddTransItems(ctx context.Context, ti *[]models.TransactionsItems) (*[]models.TransactionsItems, error) {
	if err := r.DB.WithContext(ctx).Create(&ti).Error; err != nil {
		return nil, err
	}
	return ti, nil
}
