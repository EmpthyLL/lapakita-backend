//go:build wireinject
// +build wireinject

package main

import (
	"lapakita-backend/config"
	"lapakita-backend/pkg/logger"

	"github.com/google/wire"
)

func InitializeServer() (*Server, error) {
	wire.Build(
		config.LoadConfig,
		logger.NewLogger,
		NewServer,
	)
	return &Server{}, nil
}
