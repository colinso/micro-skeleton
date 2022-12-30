package handlers

import "net/http"

type GetItem struct {
}

func NewGetItemHandler() *GetItem {
	return &GetItem{}
}

func (g GetItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Colin"))
}
