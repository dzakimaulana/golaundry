package customers

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) CustomerRespository {
	return &repository{
		DB: db,
	}
}

func (r *repository) AddNewCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error) {
	if err := r.DB.Create(cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}

func (r *repository) GetCustomerByID(ctx context.Context, id string) (*models.Customers, error) {
	var cs models.Customers
	if err := r.DB.First(&cs, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &cs, nil
}

func (r *repository) FindCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error) {
	err := r.DB.Where(&models.Customers{PhoneNumber: cs.PhoneNumber}).First(&cs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res, err := r.AddNewCustomer(ctx, cs)
			if err != nil {
				return nil, err
			}
			return res, nil
		}
		return nil, err
	}
	return cs, nil
}