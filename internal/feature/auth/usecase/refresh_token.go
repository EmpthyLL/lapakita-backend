package usecase

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"
	"lapakita-backend/pkg/jwt"
)

func (u *AuthUsecase) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.AuthResponseData, error) {
	parsedToken, err := u.JWTService.ValidateToken(req.RefreshToken)
	if err != nil || !parsedToken.Valid {
		return nil, errors.New(string(i18n.KeyRefreshTokenExpired))
	}
	claims, ok := parsedToken.Claims.(*jwt.JWTCustomClaims)
	if !ok || claims.TokenType != "refresh" {
		return nil, errors.New(string(i18n.KeyRefreshTokenExpired))
	}
	userID, err := api.DecodeRequest(claims.UserPayload.ID)
	if err != nil {
		return nil, errors.New(string(i18n.KeyRefreshTokenExpired))
	}
	user, err := u.Repo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}
	return BuildAuthResponse(u.JWTService, user)
}
