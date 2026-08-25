package middleware

import (
	"net/http"
	"strings"

	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"
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
		if err != nil || !token.Valid {
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

		c.Set(CtxUserIDKey, claims.UserID)
		c.Set(CtxUserRoleKey, claims.ActiveRole)

		c.Next()
	}
}
