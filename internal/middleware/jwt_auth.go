package middleware

import (
	"net/http"
	"strings"

	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	CtxUserIDKey        = "userID"
	CtxUserRoleKey      = "userRole"
)

func JWTAuthMiddleware(jwtService jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("lang")
		if lang == "" {
			lang = "en"
		}

		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			msg := "Authorization header not found"
			if lang == "id" {
				msg = "Header otorisasi tidak ditemukan"
			}
			api.Error(c, http.StatusUnauthorized, msg)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			msg := "Token format must be 'Bearer <token>'"
			if lang == "id" {
				msg = "Format token harus 'Bearer <token>'"
			}
			api.Error(c, http.StatusUnauthorized, msg)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			msg := "Invalid or expired token"
			if lang == "id" {
				msg = "Token tidak valid atau sudah kadaluwarsa"
			}
			api.Error(c, http.StatusUnauthorized, msg)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.JWTCustomClaims)
		if !ok {
			msg := "Failed to process token claim data"
			if lang == "id" {
				msg = "Gagal memproses data klaim token"
			}
			api.Error(c, http.StatusUnauthorized, msg)
			c.Abort()
			return
		}

		c.Set(CtxUserIDKey, claims.UserID)
		c.Set(CtxUserRoleKey, claims.Role)

		c.Next()
	}
}
