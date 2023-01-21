package grpc

import (
	"Personal/micro-skeleton/internal/models"
	context "context"
)

type CreateItemCommand interface {
	CreateItem(name string, price float32) (models.Item, error)
	GetItemById(id string) models.Item
}

type server struct {
	UnimplementedMicroSkeletonServer
	inventory CreateItemCommand
}

func (s server) CreateItem(ctx context.Context, i *Item) (*Item, error) {
	item, err := s.inventory.CreateItem(i.Name, i.Price)
	return &Item{
		Name:  item.Name,
		Price: item.Price,
	}, err
}

func (s server) GetItem(ctx context.Context, id *ItemID) (*Item, error) {
	item := s.inventory.GetItemById(id.ID)
	return &Item{
		Name:  item.Name,
		Price: item.Price,
	}, nil
}
