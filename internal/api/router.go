package api

import (
	"Personal/micro-skeleton/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(getItem *handlers.GetItem, createItem *handlers.CreateItem, updateItem *handlers.UpdateItem) *chi.Mux {
	r := chi.NewRouter()
	r.Method(http.MethodGet, "/items/{itemId}", getItem)
	r.Method(http.MethodPost, "/items", createItem)
	r.Method(http.MethodPatch, "/items/{itemId}", updateItem)
	return r
}
