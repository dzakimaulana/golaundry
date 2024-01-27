package transactions

import (
	"context"
	"time"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
)

type service struct {
	TransactionRepository
	timeout time.Duration
}

func NewService(r TransactionRepository) TransactionService {
	return &service{
		TransactionRepository: r,
		timeout:               time.Duration(2) * time.Second,
	}
}

func (s *service) AddTs(ctx context.Context, t *TransactionReq) (*TransactionRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	cs := &models.Customers{
		Name:        t.Customer.Name,
		Address:     t.Customer.Address,
		PhoneNumber: t.Customer.PhoneNumber,
	}

	csRes, err := s.FindCustomer(ctx, cs)
	if err != nil {
		return nil, err
	}
	var transitems []models.TransactionsItems
	var totalPrice int64
	var finalDuration int64
	var pricePerItems int64
	for i, ti := range t.Items {
		item, err := s.GetItemByID(ctx, ti.ItemsID.String())
		if err != nil {
			return nil, err
		}
		pricePerItems = int64(ti.Quantity * float64(item.Price))
		// append
		transitems = append(transitems, models.TransactionsItems{
			ItemsID:  t.Items[i].ItemsID,
			Quantity: t.Items[i].Quantity,
			Amount:   pricePerItems,
		})
		totalPrice += pricePerItems
		if item.Duration > finalDuration {
			finalDuration = item.Duration
		}
	}

	ts := &models.Transactions{
		ID:         uuid.New(),
		CustomerID: csRes.ID,
		UserID:     t.UserID,
		TimeIn:     time.Now().Unix(),
		TimeOut:    time.Now().Add(time.Duration(finalDuration) * time.Hour).Unix(),
		Total:      totalPrice,
	}
	r, err := s.TransactionRepository.AddTs(ctx, ts)
	if err != nil {
		return nil, err
	}

	for i := range transitems {
		transitems[i].TransactionsID = r.ID
	}

	_, err = s.AddTransItems(ctx, &transitems)
	if err != nil {
		return nil, err
	}

	res := &TransactionRes{
		ID:         r.ID,
		CustomerID: r.CustomerID,
		UserID:     r.UserID,
		TimeIn:     r.TimeIn,
		TimeOut:    r.TimeOut,
		Total:      r.Total,
	}
	return res, nil
}

func (s *service) GetAllTs(ctx context.Context) (*GetAllRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.TransactionRepository.GetAllTs(ctx)
	if err != nil {
		return nil, err
	}

	var res GetAllRes
	res.Transactions = &[]TransactionRes{}
	for _, ts := range *r {
		*res.Transactions = append(*res.Transactions, TransactionRes{
			ID:         ts.ID,
			CustomerID: ts.CustomerID,
			UserID:     ts.UserID,
			TimeIn:     ts.TimeIn,
			TimeOut:    ts.TimeOut,
			Total:      ts.Total,
		})
	}
	return &res, nil
}

func (s *service) GetTsById(ctx context.Context, id string) (*TransactionByID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.TransactionRepository.GetTsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &TransactionByID{
		ID:       r.ID,
		Customer: *r.Customer,
		Items:    *r.Items,
		TimeIn:   r.TimeIn,
		TimeOut:  r.TimeOut,
		Total:    r.Total,
	}
	return res, nil
}
