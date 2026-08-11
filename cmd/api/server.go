package main

import (
	"lapakita-backend/config"
	areaHandler "lapakita-backend/internal/feature/area/handler"
	"lapakita-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Handlers struct {
	AreaHandler *areaHandler.AreaHandler
}

func NewHandlers(areaH *areaHandler.AreaHandler) *Handlers {
	return &Handlers{
		AreaHandler: areaH,
	}
}

type Server struct {
	router   *gin.Engine
	cfg      *config.Config
	logger   *zap.Logger
	handlers *Handlers
}

func NewServer(cfg *config.Config, logger *zap.Logger, h *Handlers) *Server {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimiterMiddleware(rate.Limit(5), 10))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	apiGroup := r.Group("/api/v1")
	{
		apiGroup.GET("/areas", h.AreaHandler.SearchArea)
	}

	return &Server{
		router:   r,
		cfg:      cfg,
		logger:   logger,
		handlers: h,
	}
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.AppPort)
}
