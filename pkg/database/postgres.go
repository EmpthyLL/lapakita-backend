package database

import (
	"context"
	"fmt"
	"lapakita-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type contextKey string

const txKey contextKey = "gorm_db_trx"

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok && tx != nil {
		return tx.WithContext(ctx)
	}
	return defaultDB.WithContext(ctx)
}

func WithTransaction(ctx context.Context, db *gorm.DB, fn func(txCtx context.Context) error) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txKey, tx)
		return fn(txCtx)
	})
}

// RunInTransaction executes fn in the transaction carried by ctx when present,
// or creates a transaction for this operation otherwise.
func RunInTransaction(ctx context.Context, db *gorm.DB, fn func(tx *gorm.DB) error) error {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok && tx != nil {
		return fn(tx.WithContext(ctx))
	}
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx.WithContext(ctx))
	})
}
