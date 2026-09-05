//go:build wireinject
// +build wireinject

package main

import (
	"lapakita-backend/config"
	"lapakita-backend/internal/feature/area"
	"lapakita-backend/internal/feature/auth"
	businesstype "lapakita-backend/internal/feature/business_type"
	"lapakita-backend/internal/feature/public"
	"lapakita-backend/internal/feature/stall"
	"lapakita-backend/pkg/cache"
	"lapakita-backend/pkg/database"
	"lapakita-backend/pkg/firebase"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/logger"
	"lapakita-backend/pkg/mailer"
	"lapakita-backend/pkg/payment"
	"lapakita-backend/pkg/storage"

	"github.com/google/wire"
)

var infrastructureSet = wire.NewSet(
	config.LoadConfig,
	mailer.NewMailer,
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
	public.PublicFeatureSet,
	businesstype.BusinessTypeFeatureSet,
	auth.AuthFeatureSet,
	stall.StallFeatureSet,
)

var HandlersSet = wire.NewSet(
	wire.Struct(new(Handlers), "*"),
)

func InitializeServer() (*Server, error) {
	wire.Build(
		infrastructureSet,
		featureSet,
		HandlersSet,
		NewServer,
	)
	return &Server{}, nil
}
