package main

import (
	"lapakita-backend/config"
	areaHandler "lapakita-backend/internal/feature/area/handler"
	authHandler "lapakita-backend/internal/feature/auth/handler"
	businessTypeHandler "lapakita-backend/internal/feature/business_type/handler"
	publicHandler "lapakita-backend/internal/feature/public/handler"
	stallHandler "lapakita-backend/internal/feature/stall/handler"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Handlers struct {
	AreaHandler         *areaHandler.AreaHandler
	PublicHandler       *publicHandler.PublicHandler
	BusinessTypeHandler *businessTypeHandler.BusinessTypeHandler
	AuthHandler         *authHandler.AuthHandler
	StallHandler        *stallHandler.StallHandler
}

type Server struct {
	router     *gin.Engine
	cfg        *config.Config
	logger     *logger.Logger
	handlers   *Handlers
	jwtService *jwt.JWTService
}

func NewServer(cfg *config.Config, logger *logger.Logger, h *Handlers, jwtService *jwt.JWTService) *Server {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware(cfg.FrontendOrigin))
	r.Use(middleware.RateLimiterMiddleware(rate.Limit(5), 10))
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	apiGroup := r.Group("/api/v1")
	{
		// ---------------------------------------------------------------------
		// 1. AUTH ROUTES
		// ---------------------------------------------------------------------
		authGroup := apiGroup.Group("/auth")
		{
			// Public Auth Endpoints
			authGroup.POST("/register", h.AuthHandler.Register)
			authGroup.POST("/login", h.AuthHandler.Login)
			authGroup.POST("/google", h.AuthHandler.GoogleAuth)
			authGroup.POST("/otp/send", h.AuthHandler.SendOTP)
			authGroup.POST("/otp/verify", h.AuthHandler.VerifyOTP)
			authGroup.POST("/reset-password/:email", h.AuthHandler.ResetPassword)
			authGroup.POST("/refresh", h.AuthHandler.RefreshToken)

			// Protected Auth Endpoints (Requires Access Token)
			protectedAuth := authGroup.Group("")
			protectedAuth.Use(middleware.JWTAuthMiddleware(jwtService))
			{
				protectedAuth.PUT("/complete-profile", h.AuthHandler.CompleteProfile)
			}
		}

		// ---------------------------------------------------------------------
		// 2. AREA ROUTES
		// ---------------------------------------------------------------------
		areaGroup := apiGroup.Group("/areas")
		{
			// Public Area Search Endpoints (Tanpa Middleware Auth)
			areaGroup.GET("", h.AreaHandler.SearchGeneral)
			areaGroup.GET("/detail", h.AreaHandler.SearchDetail)

			// History Endpoints (Khusus yang butuh Optional Auth/Device ID)
			historyGroup := areaGroup.Group("/history")
			historyGroup.Use(middleware.OptionalJWTAuthMiddleware(jwtService))
			{
				historyGroup.GET("", h.AreaHandler.GetHistory)
				historyGroup.POST("", h.AreaHandler.SaveHistory)
				historyGroup.DELETE("", h.AreaHandler.ClearHistory)
				historyGroup.DELETE("/item", h.AreaHandler.DeleteItemHistory)
			}
		}

		// ---------------------------------------------------------------------
		// 3. Public ROUTES
		// ---------------------------------------------------------------------
		publicGroup := apiGroup.Group("/public")
		{
			publicGroup.GET("/faqs/:role_type", h.PublicHandler.GetFAQs)
			publicGroup.GET("/legals/:doc_type", h.PublicHandler.GetLegalDocument)
			publicGroup.POST("/contact", h.PublicHandler.SubmitContactInquiry)
		}

		// ---------------------------------------------------------------------
		// 4. BUSINESS TYPE ROUTES
		// ---------------------------------------------------------------------
		businessTypeGroup := apiGroup.Group("/business-types")
		{
			businessTypeGroup.GET("", h.BusinessTypeHandler.GetBusinessTypes)
		}

		stallGroup := apiGroup.Group("/stalls")
		{
			// Public Routes
			stallGroup.GET("", h.StallHandler.Search)
			stallGroup.GET("/:id", h.StallHandler.GetByID)
			stallGroup.GET("/:id/similar", h.StallHandler.GetSimilar)

			// Protected Routes (Owner Only)
			protectedStalls := stallGroup.Group("")
			protectedStalls.Use(middleware.JWTAuthMiddleware(jwtService))
			{
				protectedStalls.GET("/my-stalls", h.StallHandler.GetByOwner)
				protectedStalls.POST("", h.StallHandler.Create)
				protectedStalls.PUT("/:id", h.StallHandler.Update)
				protectedStalls.DELETE("/:id", h.StallHandler.Delete)
			}
		}
	}

	return &Server{
		router:     r,
		cfg:        cfg,
		logger:     logger,
		handlers:   h,
		jwtService: jwtService,
	}
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.AppPort)
}
