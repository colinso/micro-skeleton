package grpc

import (
	"Personal/micro-skeleton/internal/commands"
	"Personal/micro-skeleton/internal/models"
	context "context"
)

type CreateItemCommand interface {
	CreateItem(name string, price float32) (models.Item, error)
	GetItemById(id string) models.Item
}

type Server struct {
	UnimplementedMicroSkeletonServer
	inventory CreateItemCommand
}

func NewServer(inventory *commands.InventoryManager) *Server {
	return &Server{
		inventory: inventory,
	}
}

func (s Server) CreateItem(ctx context.Context, i *Item) (*Item, error) {
	item, err := s.inventory.CreateItem(i.Name, i.Price)
	count := int32(item.Count)
	return &Item{
		ID:       &item.ID,
		Name:     item.Name,
		Price:    item.Price,
		Count:    &count,
		ImageURL: &item.ImageURL,
	}, err
}

func (s Server) GetItem(ctx context.Context, id *ItemID) (*Item, error) {
	item := s.inventory.GetItemById(id.ID)
	count := int32(item.Count)
	return &Item{
		ID:       &item.ID,
		Name:     item.Name,
		Price:    item.Price,
		Count:    &count,
		ImageURL: &item.ImageURL,
	}, nil
}
