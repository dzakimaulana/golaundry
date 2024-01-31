package customers

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
)

type CustomerReq struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type CustomerRepository interface {
	AddNewCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error)
	GetAllCustomers(ctx context.Context) (*[]models.Customers, error)
	GetCustomerByID(ctx context.Context, id string) (*models.Customers, error)
	FindCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error)
}

type CustomerService interface {
	AddNewCustomer(ctx context.Context, cs *CustomerReq) (*models.CustomerRes, error)
	GetCustomerByID(ctx context.Context, id string) (*models.CustomerResByID, error)
	GetAllCustomers(ctx context.Context) (*[]models.CustomerRes, error)
}
