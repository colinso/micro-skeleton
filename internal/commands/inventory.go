package commands

import (
	"Personal/micro-skeleton/internal/models"
	"Personal/micro-skeleton/internal/repo"
	"context"

	"github.com/apex/log"
	"github.com/rs/xid"
)

type InventoryManager struct {
	repo repo.ItemsRepo
}

func NewInventoryManager(repo *repo.ItemsRepo) *InventoryManager {
	return &InventoryManager{
		repo: *repo,
	}
}

func (i InventoryManager) GetItemById(id string) (models.Item, error) {
	return i.repo.GetItemById(context.Background(), id)
}

func (i InventoryManager) CreateItem(name string, price float32) (models.Item, error) {
	newItem := models.Item{
		ID:    xid.New().String(),
		Name:  name,
		Price: price,
		Count: 0,
	}
	err := i.repo.CreateItem(newItem)
	if err != nil {
		log.WithError(err).Error("Error creating item")
		return models.Item{}, err
	}
	return newItem, nil
}

func (i InventoryManager) UpdateItemById(item models.Item) models.Item {
	return models.Item{}
}
