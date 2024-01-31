package items

import (
	"context"
	"time"

	"github.com/dzakimaulana/golaundry/pkg/models"
)

type Service struct {
	ItemRepository
	timeout time.Duration
}

func NewService(r ItemRepository) ItemService {
	return &Service{
		ItemRepository: r,
		timeout:        time.Duration(2) * time.Second,
	}
}

func (s *Service) AddItem(ctx context.Context, req *ItemReq) (*models.ItemRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	i := &models.Items{
		Name:     req.Name,
		Price:    req.Price,
		Unit:     req.Unit,
		Duration: req.Duration,
	}

	r, err := s.ItemRepository.AddItem(ctx, i)
	if err != nil {
		return nil, err
	}

	res := &models.ItemRes{
		ID:       r.ID,
		Name:     r.Name,
		Price:    r.Price,
		Unit:     r.Unit,
		Duration: r.Duration,
	}
	return res, nil
}

func (s *Service) GetItem(ctx context.Context, name string) (*[]models.ItemRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.ItemRepository.GetItem(ctx, name)
	if err != nil {
		return nil, err
	}

	var res []models.ItemRes
	for _, it := range *r {
		res = append(res, models.ItemRes{
			ID:       it.ID,
			Name:     it.Name,
			Price:    it.Price,
			Unit:     it.Unit,
			Duration: it.Duration,
		})
	}
	return &res, nil
}

func (s *Service) GetItemByID(ctx context.Context, id string) (*models.ItemResByID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.ItemRepository.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var resTs []models.TransItemsResIt
	for _, dataTs := range *r.Transactions {
		resTs = append(resTs, models.TransItemsResIt{
			TransactionsID: dataTs.TransactionsID,
			Quantity:       dataTs.Quantity,
			Amount:         dataTs.Amount,
		})
	}

	res := &models.ItemResByID{
		ID:           r.ID,
		Name:         r.Name,
		Price:        r.Price,
		Unit:         r.Unit,
		Duration:     r.Duration,
		Transactions: resTs,
	}
	return res, nil
}

func (s *Service) UpdateItem(ctx context.Context, req *UpdateReq) (*models.ItemRes, error) {
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

	res := &models.ItemRes{
		ID:       r.ID,
		Name:     r.Name,
		Price:    r.Price,
		Unit:     r.Unit,
		Duration: r.Duration,
	}
	return res, nil
}

func (s *Service) DeleteItem(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.ItemRepository.DeleteItem(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
