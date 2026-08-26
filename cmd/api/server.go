package main

import (
	"lapakita-backend/config"
	areaHandler "lapakita-backend/internal/feature/area/handler"
	authHandler "lapakita-backend/internal/feature/auth/handler"
	businessTypeHandler "lapakita-backend/internal/feature/business_type/handler"
	cmsHandler "lapakita-backend/internal/feature/cms/handler"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Handlers struct {
	AreaHandler         *areaHandler.AreaHandler
	CMSHandler          *cmsHandler.CMSHandler
	BusinessTypeHandler *businessTypeHandler.BusinessTypeHandler
	AuthHandler         *authHandler.AuthHandler
}

type Server struct {
	router   *gin.Engine
	cfg      *config.Config
	logger   *logger.Logger
	handlers *Handlers
}

func NewServer(cfg *config.Config, logger *logger.Logger, h *Handlers) *Server {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware(cfg.FrontendOrigin))
	r.Use(middleware.RateLimiterMiddleware(rate.Limit(5), 10))
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	apiGroup := r.Group("/api/v1")
	{

		apiGroup.GET("/areas", h.AreaHandler.SearchGeneral)

		apiGroup.GET("/areas/detail", h.AreaHandler.SearchDetail)

		apiGroup.GET("/faqs/:role_type", h.CMSHandler.GetFAQs)

		apiGroup.GET("/legals/:doc_type", h.CMSHandler.GetLegalDocument)

		apiGroup.GET("/business-types", h.BusinessTypeHandler.GetBusinessTypes)

		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/register", h.AuthHandler.Register)
			authGroup.POST("/login", h.AuthHandler.Login)
			authGroup.POST("/google", h.AuthHandler.GoogleAuth)
			authGroup.POST("/complete-profile", h.AuthHandler.CompleteProfile)
			authGroup.POST("/otp/send", h.AuthHandler.SendOTP)
			authGroup.POST("/otp/verify", h.AuthHandler.VerifyOTP)
			authGroup.POST("/reset-password", h.AuthHandler.ResetPassword)
			authGroup.POST("/refresh", h.AuthHandler.RefreshToken)
		}
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
