package customers

import (
	"context"
	"time"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
)

type service struct {
	CustomerRespository
	timeout time.Duration
}

func NewService(r CustomerRespository) CustomerService {
	return &service{
		CustomerRespository: r,
		timeout:             time.Duration(2) * time.Second,
	}
}

func (s *service) AddNewCustomer(ctx context.Context, cs *CustomerReq) (*CustomerRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	c := &models.Customers{
		ID:          uuid.New(),
		Name:        cs.Name,
		Address:     cs.Address,
		PhoneNumber: cs.PhoneNumber,
	}

	r, err := s.CustomerRespository.AddNewCustomer(ctx, c)
	if err != nil {
		return nil, err
	}

	res := &CustomerRes{
		ID:          r.ID,
		Name:        r.Name,
		Address:     r.Address,
		PhoneNumber: r.PhoneNumber,
	}
	return res, nil
}

func (s *service) GetCustomerByID(ctx context.Context, cs *CusIdReq) (*CusIdRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.CustomerRespository.GetCustomerByID(ctx, cs.ID.String())
	if err != nil {
		return nil, err
	}

	res := &CusIdRes{
		ID:          r.ID,
		Name:        r.Name,
		Address:     r.Address,
		PhoneNumber: r.PhoneNumber,
	}
	return res, nil
}
