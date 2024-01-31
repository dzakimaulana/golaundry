package items

import (
	"context"
	"testing"
	"time"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var itemsRepo = &ItemRepoMock{Mock: mock.Mock{}}
var itemsSvc = Service{
	ItemRepository: itemsRepo,
	timeout:        time.Duration(2) * time.Second,
}

func TestAddItem(t *testing.T) {

	itemReq := &ItemReq{
		Name:     "Cuci Sepatu",
		Price:    10000,
		Unit:     "pcs",
		Duration: 48,
	}

	item := &models.Items{
		Name:     itemReq.Name,
		Price:    itemReq.Price,
		Unit:     itemReq.Unit,
		Duration: itemReq.Duration,
	}
	itemsRepo.Mock.On("AddItem", mock.AnythingOfType("*context.timerCtx"), item).Return(*item, nil)
	itemResult, err := itemsSvc.AddItem(context.TODO(), itemReq)
	assert.Nil(t, err)
	assert.NotNil(t, itemResult)
	assert.Equal(t, item.ID, itemResult.ID)
}

func TestGetItem(t *testing.T) {
	name := "Cuci"
	var itemOutput []models.Items

	item := models.Items{Name: name}
	itemOutput = append(itemOutput, item)
	itemsRepo.Mock.On("GetItem", mock.AnythingOfType("*context.timerCtx"), name).Return(itemOutput, nil)

	itemResult, err := itemsSvc.GetItem(context.TODO(), name)
	assert.Nil(t, err)
	assert.NotNil(t, itemResult)
}

func TestGetItemByID(t *testing.T) {
	id := uuid.New()

	itemOutput := &models.Items{
		ID:       id,
		Name:     "Cuci Sepati",
		Price:    10000,
		Unit:     "pcs",
		Duration: 48,
		Transactions: &[]models.TransactionsItems{
			{
				TransactionsID: uuid.New(),
				ItemsID:        uuid.New(),
				Quantity:       5.4,
				Amount:         50000,
			},
		},
	}
	itemsRepo.Mock.On("GetItemByID", mock.AnythingOfType("*context.timerCtx"), id.String()).Return(*itemOutput, nil)

	itemResult, err := itemsSvc.GetItemByID(context.TODO(), id.String())
	assert.Nil(t, err)
	assert.NotNil(t, itemResult)
	assert.Equal(t, itemResult.ID, itemOutput.ID)
}

func TestUpdateItem(t *testing.T) {
	itemReq := &UpdateReq{
		ID:       uuid.New(),
		Name:     "Cuci Sepati",
		Price:    10000,
		Unit:     "pcs",
		Duration: 48,
	}

	item := &models.Items{
		ID:       itemReq.ID,
		Name:     itemReq.Name,
		Price:    itemReq.Price,
		Unit:     itemReq.Unit,
		Duration: itemReq.Duration,
	}
	itemsRepo.Mock.On("UpdateItem", mock.AnythingOfType("*context.timerCtx"), item).Return(*item, nil)

	itemResult, err := itemsSvc.UpdateItem(context.TODO(), itemReq)
	assert.Nil(t, err)
	assert.NotNil(t, itemResult)
}

func TestDeleteItem(t *testing.T) {
	id := uuid.New()
	itemsRepo.Mock.On("DeleteItem", mock.AnythingOfType("*context.timerCtx"), id.String()).Return(nil)

	err := itemsSvc.DeleteItem(context.TODO(), id.String())
	assert.Nil(t, err)
}
