package customers

import (
	"context"
	"time"

	"github.com/dzakimaulana/golaundry/pkg/models"
)

type Service struct {
	CustomerRepository
	timeout time.Duration
}

func NewService(r CustomerRepository) CustomerService {
	return &Service{
		CustomerRepository: r,
		timeout:            time.Duration(2) * time.Second,
	}
}

func (s *Service) AddNewCustomer(ctx context.Context, cs *CustomerReq) (*models.CustomerRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	c := &models.Customers{
		Name:        cs.Name,
		Address:     cs.Address,
		PhoneNumber: cs.PhoneNumber,
	}

	r, err := s.CustomerRepository.AddNewCustomer(ctx, c)
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

func (s *Service) GetAllCustomers(ctx context.Context) (*[]models.CustomerRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.CustomerRepository.GetAllCustomers(ctx)
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

func (s *Service) GetCustomerByID(ctx context.Context, id string) (*models.CustomerResByID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.CustomerRepository.GetCustomerByID(ctx, id)
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
