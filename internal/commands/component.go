package commands

import "github.com/google/wire"

var Component = wire.NewSet(NewInventoryManager)
