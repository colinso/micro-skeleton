package handlers

import "github.com/google/wire"

var Component = wire.NewSet(NewGetItemHandler, NewCreateItemHandler, NewUpdateItemHandler)
