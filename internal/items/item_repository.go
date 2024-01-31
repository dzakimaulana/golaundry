package items

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) ItemRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) AddItem(ctx context.Context, item *models.Items) (*models.Items, error) {
	item.ID = uuid.New()
	if err := r.DB.WithContext(ctx).Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *repository) GetItem(ctx context.Context, name string) (*[]models.Items, error) {
	var items []models.Items
	if name == "" {
		if err := r.DB.WithContext(ctx).Find(&items).Error; err != nil {
			return nil, err
		}
		return &items, nil
	}
	if err := r.DB.WithContext(ctx).Where("name = ?", name).Find(&items).Error; err != nil {
		return nil, err
	}
	return &items, nil
}

func (r *repository) GetItemByID(ctx context.Context, id string) (*models.Items, error) {
	var item models.Items
	if err := r.DB.WithContext(ctx).Model(&models.Items{}).Preload("Transactions").First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *repository) UpdateItem(ctx context.Context, item *models.Items) (*models.Items, error) {
	if err := r.DB.WithContext(ctx).Updates(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *repository) DeleteItem(ctx context.Context, id string) error {
	if err := r.DB.WithContext(ctx).Delete(&models.Items{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
