package handlers

import (
	"Personal/micro-skeleton/internal/commands"
	"Personal/micro-skeleton/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

type GetItemCommand interface {
	GetItemById(id string) (models.Item, error)
}

type GetItem struct {
	inventory GetItemCommand
}

func NewGetItemHandler(inventory *commands.InventoryManager) *GetItem {
	return &GetItem{
		inventory: inventory,
	}
}

func (g GetItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlSplit := strings.Split(r.URL.Path, "/")
	id := urlSplit[len(urlSplit)-1]

	item, err := g.inventory.GetItemById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, _ := json.Marshal(item)
	w.Write(b)
}
