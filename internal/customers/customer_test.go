package customers

import (
	"context"
	"testing"
	"time"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var customerRepo = &CustomerRepoMock{Mock: mock.Mock{}}
var customerSvc = Service{
	CustomerRepository: customerRepo,
	timeout:            time.Duration(2) * time.Second,
}

func TestAddNewCustomer(t *testing.T) {
	customerReq := &CustomerReq{
		Name:        "Budi",
		Address:     "Jl. Keranggan",
		PhoneNumber: "086734556734",
	}

	customer := &models.Customers{
		Name:         customerReq.Name,
		Address:      customerReq.Address,
		PhoneNumber:  customerReq.PhoneNumber,
		Transactions: nil,
	}
	customerRepo.Mock.On("AddNewCustomer", mock.AnythingOfType("*context.timerCtx"), customer).Return(*customer, nil)

	createdCustomer, err := customerSvc.AddNewCustomer(context.TODO(), customerReq)
	assert.Nil(t, err)
	assert.NotNil(t, createdCustomer)
	assert.Equal(t, customer.ID, createdCustomer.ID)
}

func TestGetAllCustomers(t *testing.T) {
	var alluser []models.Customers
	var expOutput []models.CustomerRes
	customerRepo.Mock.On("GetAllCustomers", mock.AnythingOfType("*context.timerCtx")).Return(alluser, nil)
	getAllCus, err := customerSvc.GetAllCustomers(context.TODO())
	allCusRes := *getAllCus
	assert.Nil(t, err)
	assert.NotNil(t, getAllCus)
	assert.Equal(t, expOutput, allCusRes)
}

func TestGetCustomerByID(t *testing.T) {
	// Sample data for the mock response
	id := uuid.New()

	customer := &models.Customers{
		ID:          id,
		Name:        "Joko",
		Address:     "Jl. Oke",
		PhoneNumber: "09877635521",
		Transactions: &[]models.Transactions{
			{
				ID:         uuid.New(),
				CustomerID: id,
				UserID:     uuid.New(),
				TimeIn:     48,
				TimeOut:    72,
				Total:      11000,
			},
		},
	}

	customerRepo.Mock.On("GetCustomerByID", mock.AnythingOfType("*context.timerCtx"), id.String()).Return(*customer, nil)

	oneCustomer, err := customerSvc.GetCustomerByID(context.TODO(), id.String())
	assert.Nil(t, err)
	assert.NotNil(t, oneCustomer)
	assert.Equal(t, customer.ID, oneCustomer.ID)
}
