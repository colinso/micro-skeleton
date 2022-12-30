package repo

import "github.com/google/wire"

var Component = wire.NewSet(NewItemsRepo)
