//go:build wireinject
// +build wireinject

package main

import (
	"lapakita-backend/config"
	"lapakita-backend/internal/feature/area"
	"lapakita-backend/pkg/cache"
	"lapakita-backend/pkg/database"
	"lapakita-backend/pkg/firebase"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/logger"
	"lapakita-backend/pkg/payment"
	"lapakita-backend/pkg/storage"

	"github.com/google/wire"
)

var infrastructureSet = wire.NewSet(
	config.LoadConfig,
	logger.NewLogger,
	database.NewPostgresDB,
	cache.NewRedisClient,
	jwt.NewJWTService,
	firebase.NewFirebaseService,
	storage.NewImageKitService,
	payment.NewPaymentService,
)

var featureSet = wire.NewSet(
	area.AreaFeatureSet,
)

func InitializeServer() (*Server, error) {
	wire.Build(
		infrastructureSet,
		featureSet,
		NewHandlers,
		NewServer,
	)
	return &Server{}, nil
}
