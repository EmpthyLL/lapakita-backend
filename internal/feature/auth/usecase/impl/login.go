package impl

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/auth/dto"
	authmessages "lapakita-backend/internal/feature/auth/messages"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/pkg/i18n"

	"golang.org/x/crypto/bcrypt"
)

func Login(ctx context.Context, u *usecase.AuthUsecase, req dto.LoginRequest) (*dto.AuthResponseData, error) {
	user, err := u.Repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New(string(authmessages.KeyWrongPassword))
	}
	return usecase.BuildAuthResponse(u.JWTService, user)
}
