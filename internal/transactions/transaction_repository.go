package transactions

import (
	"context"

	"github.com/dzakimaulana/golaundry/internal/customers"
	"github.com/dzakimaulana/golaundry/internal/items"
	"github.com/dzakimaulana/golaundry/internal/transitems"
	"github.com/dzakimaulana/golaundry/pkg/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) TransactionRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) AddTs(ctx context.Context, t *models.Transactions) (*models.Transactions, error) {
	if err := r.DB.WithContext(ctx).Create(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (r *repository) GetAllTs(ctx context.Context) (*[]models.Transactions, error) {
	var ts []models.Transactions
	if err := r.DB.WithContext(ctx).Find(&ts).Error; err != nil {
		return nil, err
	}
	return &ts, nil
}

func (r *repository) GetTsByID(ctx context.Context, id string) (*models.Transactions, error) {
	var t models.Transactions
	if err := r.DB.WithContext(ctx).Model(&models.Transactions{}).Preload("Items").Preload("User").Preload("Customer").First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *repository) FindCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error) {
	res, err := customers.NewRepository(r.DB).FindCustomer(ctx, cs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) GetItemByID(ctx context.Context, id string) (*models.Items, error) {
	res, err := items.NewRepository(r.DB).GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) AddTransItems(ctx context.Context, ti *[]models.TransactionsItems) (*[]models.TransactionsItems, error) {
	res, err := transitems.NewRepository(r.DB).AddTransItems(ctx, ti)
	if err != nil {
		return nil, err
	}
	return res, nil
}
