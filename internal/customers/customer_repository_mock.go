package customers

import (
	"context"
	"errors"
	"fmt"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/stretchr/testify/mock"
)

type CustomerRepoMock struct {
	Mock mock.Mock
}

func (r *CustomerRepoMock) AddNewCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error) {
	arguments := r.Mock.Called(ctx, cs)
	if errArg := arguments.Get(1); errArg != nil {
		err := errArg.(error)
		return nil, err
	}
	customer := arguments.Get(0).(models.Customers)
	return &customer, nil
}

func (r *CustomerRepoMock) GetAllCustomers(ctx context.Context) (*[]models.Customers, error) {
	arguments := r.Mock.Called(ctx)
	if arguments.Get(0) == nil {
		return nil, errors.New("parameter not found")
	}
	customer := arguments.Get(0).([]models.Customers)
	return &customer, nil
}

func (r *CustomerRepoMock) GetCustomerByID(ctx context.Context, id string) (*models.Customers, error) {
	arguments := r.Mock.Called(ctx, id)
	if errArg := arguments.Get(1); errArg != nil {
		err, ok := errArg.(error)
		if !ok {
			return nil, fmt.Errorf("unexpected type for error argument")
		}
		return nil, err
	}

	customer, ok := arguments.Get(0).(models.Customers)
	if !ok {
		return nil, fmt.Errorf("unexpected type for customer argument")
	}

	return &customer, nil
}

func (r *CustomerRepoMock) FindCustomer(ctx context.Context, cs *models.Customers) (*models.Customers, error) {
	arguments := r.Mock.Called(ctx, cs)
	if arguments.Get(0) == nil && arguments.Get(1) == nil {
		return nil, errors.New("parameter not found")
	}
	customer := arguments.Get(1).(models.Customers)
	return &customer, nil
}
