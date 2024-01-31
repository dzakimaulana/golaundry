package items

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/stretchr/testify/mock"
)

type ItemRepoMock struct {
	Mock mock.Mock
}

func (r *ItemRepoMock) AddItem(ctx context.Context, item *models.Items) (*models.Items, error) {
	arguments := r.Mock.Called(ctx, item)
	if errArg := arguments.Get(1); errArg != nil {
		err := errArg.(error)
		return nil, err
	}
	it := arguments.Get(0).(models.Items)
	return &it, nil
}

func (r *ItemRepoMock) GetItem(ctx context.Context, name string) (*[]models.Items, error) {
	arguments := r.Mock.Called(ctx, name)
	if errArg := arguments.Get(1); errArg != nil {
		err := errArg.(error)
		return nil, err
	}
	it := arguments.Get(0).([]models.Items)
	return &it, nil
}

func (r *ItemRepoMock) GetItemByID(ctx context.Context, id string) (*models.Items, error) {
	arguments := r.Mock.Called(ctx, id)
	if errArg := arguments.Get(1); errArg != nil {
		err := errArg.(error)
		return nil, err
	}
	it := arguments.Get(0).(models.Items)
	return &it, nil
}

func (r *ItemRepoMock) UpdateItem(ctx context.Context, item *models.Items) (*models.Items, error) {
	arguments := r.Mock.Called(ctx, item)
	if errArg := arguments.Get(1); errArg != nil {
		err := errArg.(error)
		return nil, err
	}
	it := arguments.Get(0).(models.Items)
	return &it, nil
}

func (r *ItemRepoMock) DeleteItem(ctx context.Context, id string) error {
	arguments := r.Mock.Called(ctx, id)
	if errArg := arguments.Get(0); errArg != nil {
		err := errArg.(error)
		return err
	}
	return nil
}
