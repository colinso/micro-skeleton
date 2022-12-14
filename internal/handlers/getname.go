package handlers

import "net/http"

type GetName struct {
}

func NewGetNameHandler() GetName {
	return GetName{}
}

func (g GetName) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Colin"))
}
