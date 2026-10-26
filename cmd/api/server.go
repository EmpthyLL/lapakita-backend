package main

import (
	"lapakita-backend/config"
	authHandler "lapakita-backend/internal/feature/auth/handler"
	masterHandler "lapakita-backend/internal/feature/master/handler"
	stallHandler "lapakita-backend/internal/feature/stall/handler"
	userHandler "lapakita-backend/internal/feature/user/handler"
	walletHandler "lapakita-backend/internal/feature/wallet/handler"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Handlers struct {
	MasterHandler *masterHandler.MasterHandler
	AuthHandler   *authHandler.AuthHandler
	StallHandler  *stallHandler.StallHandler
	WalletHandler *walletHandler.WalletHandler
	UserHandler   *userHandler.UserHandler
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
			areaGroup.GET("", h.MasterHandler.SearchGeneral)
			areaGroup.GET("/detail", h.MasterHandler.SearchDetail)

			// History Endpoints (Khusus yang butuh Optional Auth/Device ID)
			historyGroup := areaGroup.Group("/history")
			historyGroup.Use(middleware.OptionalJWTAuthMiddleware(jwtService))
			{
				historyGroup.GET("", h.MasterHandler.GetHistory)
				historyGroup.POST("", h.MasterHandler.SaveHistory)
				historyGroup.DELETE("", h.MasterHandler.ClearHistory)
				historyGroup.DELETE("/item", h.MasterHandler.DeleteItemHistory)
			}
		}

		// ---------------------------------------------------------------------
		// 3. Public ROUTES
		// ---------------------------------------------------------------------
		publicGroup := apiGroup.Group("/public")
		{
			publicGroup.GET("/faqs/:role_type", h.MasterHandler.GetFAQs)
			publicGroup.GET("/legals/:doc_type", h.MasterHandler.GetLegalDocument)
			publicGroup.POST("/contact", h.MasterHandler.SubmitContact)
		}

		// ---------------------------------------------------------------------
		// 4. BUSINESS TYPE ROUTES
		// ---------------------------------------------------------------------
		businessTypeGroup := apiGroup.Group("/business-types")
		{
			businessTypeGroup.GET("", h.MasterHandler.GetBusinessTypes)
		}

		// ---------------------------------------------------------------------
		// 4. STALL ROUTES
		// ---------------------------------------------------------------------
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

		// ---------------------------------------------------------------------
		// 6. USER / SETTINGS ROUTES (Protected)
		// ---------------------------------------------------------------------
		userGroup := apiGroup.Group("/users")
		userGroup.Use(middleware.JWTAuthMiddleware(jwtService))
		{
			// General Profile
			profileGroup := userGroup.Group("/profile")
			{
				profileGroup.GET("", h.UserHandler.GetGeneralProfile)
				profileGroup.PUT("", h.UserHandler.UpdateGeneralProfile)
			}

			// Phone Numbers
			phoneGroup := userGroup.Group("/phone")
			{
				phoneGroup.GET("", h.UserHandler.GetPhoneNumbers)
				phoneGroup.POST("", h.UserHandler.AddPhoneNumber)
				phoneGroup.PUT("/:index", h.UserHandler.UpdatePhoneNumber)
				phoneGroup.DELETE("/:index", h.UserHandler.DeletePhoneNumber)
			}

			// Security & Password
			passwordGroup := userGroup.Group("/password")
			{
				passwordGroup.PUT("", h.UserHandler.UpdatePassword)
			}

			// Persona Profile
			personaGroup := userGroup.Group("/persona")
			{
				personaGroup.GET("/:role", h.UserHandler.GetPersonaProfile)
				personaGroup.PUT("/:role", h.UserHandler.UpdatePersonaProfile)
			}

			// Legal Documents
			documentGroup := userGroup.Group("/document")
			{
				documentGroup.GET("", h.UserHandler.GetDocument)
				documentGroup.POST("", h.UserHandler.UploadDocument)
				documentGroup.DELETE("/:id", h.UserHandler.DeleteDocument)
			}

			// Payout Methods
			payoutMethodGroup := apiGroup.Group("/payment-method")
			payoutMethodGroup.Use(middleware.JWTAuthMiddleware(jwtService))
			{
				payoutMethodGroup.GET("", h.WalletHandler.GetPaymentMethods)
				payoutMethodGroup.POST("", h.WalletHandler.CreatePaymentMethod)
				payoutMethodGroup.DELETE("/:id", h.WalletHandler.DeletePaymentMethod)
				payoutMethodGroup.PUT("/:id/primary", h.WalletHandler.SetPrimaryPaymentMethod)
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
