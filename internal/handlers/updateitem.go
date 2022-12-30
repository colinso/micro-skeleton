package handlers

import "net/http"

type UpdateItem struct {
}

func NewUpdateItemHandler() *UpdateItem {
	return &UpdateItem{}
}

func (u UpdateItem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Colin"))
}
