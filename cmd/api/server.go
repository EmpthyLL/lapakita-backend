package main

import (
	"lapakita-backend/config"
	"lapakita-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	logger *zap.Logger
}

func NewServer(cfg *config.Config, logger *zap.Logger) *Server {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.Use(middleware.RateLimiterMiddleware(rate.Limit(5), 10))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return &Server{
		router: r,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.AppPort)
}
