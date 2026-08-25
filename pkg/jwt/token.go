package jwt

import (
	"fmt"
	"time"

	"lapakita-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID                string `json:"user_id"`
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	Phone                 string `json:"phone,omitempty"`
	AvatarURL             string `json:"avatar_url,omitempty"`
	ActiveRole            string `json:"active_role"`
	SubscriptionPlan      string `json:"subscription_plan"`
	SubscriptionExpiresAt string `json:"subscription_expires_at,omitempty"`
	TokenType             string `json:"token_type"` // "access" atau "refresh"
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTService interface {
	GenerateTokenPair(claims JWTCustomClaims) (*TokenPair, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey     []byte
	issuer        string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewJWTService(cfg *config.Config) JWTService {
	duration, err := time.ParseDuration(cfg.JWTExpiry)
	if err != nil {
		duration = 24 * time.Hour
	}

	return &jwtService{
		secretKey:     []byte(cfg.JWTSecret),
		issuer:        cfg.AppName,
		accessExpiry:  duration,
		refreshExpiry: 30 * 24 * time.Hour,
	}
}

func (j *jwtService) GenerateTokenPair(claims JWTCustomClaims) (*TokenPair, error) {
	now := time.Now()

	// 1. Generate Access Token
	accessClaims := claims
	accessClaims.TokenType = "access"
	accessClaims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(j.accessExpiry)),
		IssuedAt:  jwt.NewNumericDate(now),
		Issuer:    j.issuer,
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := accessTokenObj.SignedString(j.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	// 2. Generate Refresh Token
	refreshClaims := JWTCustomClaims{
		UserID:    claims.UserID,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    j.issuer,
		},
	}

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenObj.SignedString(j.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})
}
