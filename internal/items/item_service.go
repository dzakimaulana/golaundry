package items

import (
	"context"
	"time"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
)

type service struct {
	ItemRepository
	timeout time.Duration
}

func NewService(r ItemRepository) ItemService {
	return &service{
		ItemRepository: r,
		timeout:        time.Duration(2) * time.Second,
	}
}

func (s *service) AddItem(ctx context.Context, req *ItemReq) (*ItemRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	i := &models.Items{
		ID:       uuid.New(),
		Name:     req.Name,
		Price:    req.Price,
		Unit:     req.Unit,
		Duration: req.Duration,
	}

	r, err := s.ItemRepository.AddItem(ctx, i)
	if err != nil {
		return nil, err
	}

	res := &ItemRes{
		ID:       r.ID,
		Name:     r.Name,
		Price:    r.Price,
		Unit:     r.Unit,
		Duration: r.Duration,
	}
	return res, nil
}

func (s *service) GetItem(ctx context.Context, name string) (*GetRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.ItemRepository.GetItem(ctx, name)
	if err != nil {
		return nil, err
	}

	var res GetRes
	res.Items = &[]ItemRes{}
	for _, it := range *r {
		*res.Items = append(*res.Items, ItemRes{
			ID:       it.ID,
			Name:     it.Name,
			Price:    it.Price,
			Unit:     it.Unit,
			Duration: it.Duration,
		})
	}
	return &res, nil
}

func (s *service) GetItemByID(ctx context.Context, id string) (*ItemResByID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.ItemRepository.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &ItemResByID{
		ID:           r.ID,
		Name:         r.Name,
		Price:        r.Price,
		Unit:         r.Unit,
		Duration:     r.Duration,
		Transactions: *r.Transactions,
	}
	return res, nil
}

func (s *service) UpdateItem(ctx context.Context, req *UpdateReq) (*UpdateRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	i := &models.Items{
		ID:       req.ID,
		Name:     req.Name,
		Price:    req.Price,
		Unit:     req.Unit,
		Duration: req.Duration,
	}

	r, err := s.ItemRepository.UpdateItem(ctx, i)
	if err != nil {
		return nil, err
	}

	res := &UpdateRes{
		ID:       r.ID,
		Name:     r.Name,
		Price:    r.Price,
		Unit:     r.Unit,
		Duration: r.Duration,
	}
	return res, nil
}

func (s *service) DeleteItem(ctx context.Context, req *DelReq) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.ItemRepository.DeleteItem(ctx, req.ID.String())
	if err != nil {
		return err
	}
	return nil
}
