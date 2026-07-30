package logger

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	defer log.Sync()
	return log, nil
}
