package usecase

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) UpdatePassword(ctx context.Context, userID uuid.UUID, req dto.ChangePasswordRequest) error {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}
	if user.PasswordHash != "" {
		if req.CurrentPassword == nil || *req.CurrentPassword == "" {
			return errors.New(string(i18n.KeyUserPasswordIncorrect))
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(*req.CurrentPassword)); err != nil {
			return errors.New(string(i18n.KeyUserPasswordIncorrect))
		}
	}
	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return u.repo.UpdatePassword(ctx, userID, string(newHash))
}
