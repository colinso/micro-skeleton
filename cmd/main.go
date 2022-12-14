package main

import (
	"Personal/micro-skeleton/internal/api"
	"os"
)

func main() {
	if err := api.Start(); err != nil {
		os.Exit(1)
	}
}
