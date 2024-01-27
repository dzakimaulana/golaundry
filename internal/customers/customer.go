package customers

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
)

type CustomerReq struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type CustomerRes struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
}

type CusIdReq struct {
	ID uuid.UUID `json:"id"`
}

type CusIdRes struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
}

type CustomerRespository interface {
	AddNewCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error)
	GetCustomerByID(ctx context.Context, id string) (*models.Customers, error)
	FindCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error)
}

type CustomerService interface {
	AddNewCustomer(ctx context.Context, cs *CustomerReq) (*CustomerRes, error)
	GetCustomerByID(ctx context.Context, cs *CusIdReq) (*CusIdRes, error)
}
