package usecase

import (
	"context"
	"errors"
	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/pkg/i18n"

	"golang.org/x/crypto/bcrypt"
)

func (u *AuthUsecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponseData, error) {
	user, err := u.repo.FindUserByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New(string(i18n.KeyWrongPassword))
	}
	return usecase.BuildAuthResponse(u.jwtService, user)
}
