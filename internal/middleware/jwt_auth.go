package middleware

import (
	"net/http"
	"strings"

	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"
	"lapakita-backend/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	CtxUserIDKey        = "userID"
)

// JWTAuthMiddleware menggunakan pointer *jwt.JWTService
func JWTAuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyAuthHeaderNotFound))
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyTokenFormatInvalid))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || token == nil || !token.Valid {
			api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyTokenInvalidOrExpired))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.JWTCustomClaims)
		if !ok || claims.TokenType != "access" {
			api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyTokenClaimsFailed))
			c.Abort()
			return
		}

		// Simpan User ID ke context Gin
		c.Set(CtxUserIDKey, claims.UserPayload.ID)

		c.Next()
	}
}

// OptionalJWTAuthMiddleware mengekstrak UserID jika Authorization Bearer header valid.
func OptionalJWTAuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader != "" && strings.HasPrefix(authHeader, BearerPrefix) {
			tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
			token, err := jwtService.ValidateToken(tokenString)
			if err == nil && token != nil && token.Valid {
				if claims, ok := token.Claims.(*jwt.JWTCustomClaims); ok && claims.TokenType == "access" {
					c.Set(CtxUserIDKey, claims.UserPayload.ID)
				}
			}
		}
		c.Next()
	}
}

// GetUserIDFromContext helper function untuk mengambil userID secara aman
func GetUserIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	val, exists := c.Get(CtxUserIDKey)
	if !exists {
		return uuid.Nil, false
	}

	userIDStr, ok := val.(string)
	if !ok {
		return uuid.Nil, false
	}

	parsedUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, false
	}

	return parsedUUID, true
}
