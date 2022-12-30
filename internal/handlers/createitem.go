package handlers

import (
	"Personal/micro-skeleton/internal/commands"
	"Personal/micro-skeleton/internal/models"
	"encoding/json"
	"net/http"
)

type CreateItemCommand interface {
	CreateItem(name string, price float32) (models.Item, error)
}

type CreateItem struct {
	inventory CreateItemCommand
}

func NewCreateItemHandler(inventory *commands.InventoryManager) *CreateItem {
	return &CreateItem{
		inventory: inventory,
	}
}

func (c CreateItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body models.CreateItemRequest
	json.NewDecoder(r.Body).Decode(&body)

	item, err := c.inventory.CreateItem(body.Name, body.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, _ := json.Marshal(item)
	w.Write(b)
}
