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
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			api.ErrorResponse(c, http.StatusUnauthorized, "Header otorisasi tidak ditemukan", nil)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			api.ErrorResponse(c, http.StatusUnauthorized, "Format token harus 'Bearer <token>'", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			api.ErrorResponse(c, http.StatusUnauthorized, "Token tidak valid atau sudah kadaluwarsa", nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.JWTCustomClaims)
		if !ok {
			api.ErrorResponse(c, http.StatusUnauthorized, "Gagal memproses data klaim token", nil)
			c.Abort()
			return
		}

		// Simpan userID dan role ke konteks Gin agar bisa diakses di layer Handler
		c.Set(CtxUserIDKey, claims.UserID)
		c.Set(CtxUserRoleKey, claims.Role)

		c.Next()
	}
}