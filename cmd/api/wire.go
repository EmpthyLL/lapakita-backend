//go:build wireinject
// +build wireinject

package main

import (
	"lapakita-backend/config"
	"lapakita-backend/internal/feature/auth"
	"lapakita-backend/internal/feature/master"
	"lapakita-backend/internal/feature/stall"
	"lapakita-backend/internal/feature/user"
	"lapakita-backend/internal/feature/wallet"
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
	storage.NewSeaweedFSService,
	payment.NewPaymentService,
)

var featureSet = wire.NewSet(
	master.MasterFeatureSet,
	auth.AuthFeatureSet,
	wallet.WalletFeatureSet,
	stall.StallFeatureSet,
	user.UserFeatureSet,
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
