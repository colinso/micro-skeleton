package api

import (
	"Personal/micro-skeleton/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Method(http.MethodGet, "/name", handlers.NewGetNameHandler())
	return r
}
