package impl

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/user/dto"
	usermessages "lapakita-backend/internal/feature/user/messages"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, req dto.ChangePasswordRequest) error {
	user, err := u.Repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}
	if user.PasswordHash != "" {
		if req.CurrentPassword == nil || *req.CurrentPassword == "" {
			return errors.New(string(usermessages.KeyUserPasswordIncorrect))
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(*req.CurrentPassword)); err != nil {
			return errors.New(string(usermessages.KeyUserPasswordIncorrect))
		}
	}
	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return u.Repo.UpdatePassword(ctx, userID, string(newHash))
}
