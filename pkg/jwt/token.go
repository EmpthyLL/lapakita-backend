package jwt

import (
	"fmt"
	"time"

	"lapakita-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID string, role string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey []byte
	issuer    string
	expiry    time.Duration
}

func NewJWTService(cfg *config.Config) JWTService {
	duration, err := time.ParseDuration(cfg.JWTExpiry)
	if err != nil {
		duration = 24 * time.Hour // Default 24 jam jika format di config gagal di-parse
	}

	return &jwtService{
		secretKey: []byte(cfg.JWTSecret),
		issuer:    cfg.AppName,
		expiry:    duration,
	}
}

func (j *jwtService) GenerateToken(userID string, role string) (string, error) {
	claims := &JWTCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", fmt.Errorf("gagal menandatangani token: %w", err)
	}

	return signedToken, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak valid: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})
}