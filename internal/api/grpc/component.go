package grpc

import "github.com/google/wire"

var Component = wire.NewSet(NewServer)
