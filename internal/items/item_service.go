package items

import (
	"context"
	"time"

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

func (s *service) AddItem(ctx context.Context, req *AddReq) (*AddRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	i := &Items{
		ID:    uuid.New(),
		Name:  req.Name,
		Price: req.Price,
		Unit:  req.Unit,
		Time:  req.Time,
	}

	r, err := s.ItemRepository.AddItem(ctx, i)
	if err != nil {
		return nil, err
	}

	res := &AddRes{
		ID:    r.ID,
		Name:  r.Name,
		Price: r.Price,
		Unit:  r.Unit,
		Time:  r.Time,
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

	res := &GetRes{
		Items: r,
	}
	return res, nil
}

func (s *service) UpdateItem(ctx context.Context, req *UpdateReq) (*UpdateRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	i := &Items{
		ID:    req.ID,
		Name:  req.Name,
		Price: req.Price,
		Unit:  req.Unit,
		Time:  req.Time,
	}

	r, err := s.ItemRepository.UpdateItem(ctx, i)
	if err != nil {
		return nil, err
	}

	res := &UpdateRes{
		ID:    r.ID,
		Name:  r.Name,
		Price: r.Price,
		Unit:  r.Unit,
		Time:  r.Time,
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
