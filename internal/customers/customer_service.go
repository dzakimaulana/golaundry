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

func (s *service) AddNewCustomer(ctx context.Context, cs *CustomerReq) (*models.CustomerRes, error) {
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

	res := &models.CustomerRes{
		ID:          r.ID,
		Name:        r.Name,
		Address:     r.Address,
		PhoneNumber: r.PhoneNumber,
	}
	return res, nil
}

func (s *service) GetAllCustomers(ctx context.Context) (*[]models.CustomerRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.CustomerRespository.GetAllCustomer(ctx)
	if err != nil {
		return nil, err
	}

	var allCs []models.CustomerRes
	for _, oneCs := range *r {
		allCs = append(allCs, models.CustomerRes{
			ID:          oneCs.ID,
			Name:        oneCs.Name,
			Address:     oneCs.Address,
			PhoneNumber: oneCs.PhoneNumber,
		})
	}
	return &allCs, nil
}

func (s *service) GetCustomerByID(ctx context.Context, id string) (*models.CustomerResByID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.CustomerRespository.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var tsResponse []models.TransactionsResCus
	for _, tsData := range *r.Transactions {
		tsResponse = append(tsResponse, models.TransactionsResCus{
			ID:      tsData.ID,
			UserID:  tsData.UserID,
			TimeIn:  tsData.TimeIn,
			TimeOut: tsData.TimeOut,
			Total:   tsData.Total,
		})
	}

	res := &models.CustomerResByID{
		ID:           r.ID,
		Name:         r.Name,
		Address:      r.Address,
		PhoneNumber:  r.PhoneNumber,
		Transactions: tsResponse,
	}
	return res, nil
}
