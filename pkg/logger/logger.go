package logger

import (
	"context"
	"os"

	"lapakita-backend/config"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(cfg *config.Config) (*Logger, error) {
	// 1. Parse Log Level dari Config
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		level = zapcore.DebugLevel
	}

	// 2. Setup JSON Production Encoder Config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(zapcore.AddSync(os.Stdout)), // Gunakan os.Stdout di sini
		level,
	)

	zapLog := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zapLog.Info("Logger initialized successfully", zap.String("level", level.String()))

	return &Logger{Logger: zapLog}, nil
}

// WithContext mengekstrak TraceID & SpanID dari OpenTelemetry Context
func (l *Logger) WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return l.Logger
	}

	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		return l.Logger
	}

	return l.Logger.With(
		zap.String("trace_id", spanCtx.TraceID().String()),
		zap.String("span_id", spanCtx.SpanID().String()),
	)
}
