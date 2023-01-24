//go:build wireinject
// +build wireinject

package wire

import (
	"Personal/micro-skeleton/internal/api"
	"Personal/micro-skeleton/internal/api/grpc"
	"Personal/micro-skeleton/internal/commands"
	"Personal/micro-skeleton/internal/config"
	"Personal/micro-skeleton/internal/db"
	"Personal/micro-skeleton/internal/handlers"
	"Personal/micro-skeleton/internal/repo"

	"github.com/google/wire"
)

func ConfigureServer() *api.Server {
	wire.Build(
		config.NewConfig,
		api.NewServer,
		api.NewRouter,
		db.NewDBConnection,

		handlers.Component,
		commands.Component,
		repo.Component,
		grpc.Component,
	)
	return &api.Server{}
}
