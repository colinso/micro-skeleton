package main

import (
	"Personal/micro-skeleton/internal/wire"
	"os"

	"github.com/apex/log"
)

func main() {
	if err := wire.ConfigureServer().Start(); err != nil {
		log.WithError(err).Error("Application shut down")
		os.Exit(1)
	}
}
