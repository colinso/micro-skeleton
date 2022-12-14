//go:build wireinject
// +build wireinject

package wire

import (
	"Personal/micro-skeleton/internal/api"
	"Personal/micro-skeleton/internal/config"

	"github.com/google/wire"
)

func ConfigureServer() *api.Server {
	wire.Build(
		config.NewConfig,
		api.NewServer,
		api.NewRouter,
	)
	return &api.Server{}
}
