package items

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) ItemRepository {
	db.AutoMigrate(&Items{})
	return &repository{
		DB: db,
	}
}

func (r *repository) AddItem(ctx context.Context, item *Items) (*Items, error) {
	if err := r.DB.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *repository) GetItem(ctx context.Context, name string) (*[]Items, error) {
	var items []Items
	if name == "" {
		if err := r.DB.Find(&items).Error; err != nil {
			return nil, err
		}
		return &items, nil
	}
	if err := r.DB.Where("name = ?", name).Find(&items).Error; err != nil {
		return nil, err
	}
	return &items, nil
}

func (r *repository) UpdateItem(ctx context.Context, item *Items) (*Items, error) {
	if err := r.DB.Updates(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *repository) DeleteItem(ctx context.Context, id string) error {
	if err := r.DB.Delete(&Items{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
