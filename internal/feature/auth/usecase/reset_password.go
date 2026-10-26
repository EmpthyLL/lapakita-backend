package usecase

import (
	"context"
	"errors"
	"fmt"

	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/pkg/i18n"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

func (u *AuthUsecase) ResetPassword(ctx context.Context, email string, req dto.ResetPasswordRequest) error {
	verifyTokenKey := fmt.Sprintf("verified:reset_password:%s", email)
	storedVal, err := u.RDB.Get(ctx, verifyTokenKey).Result()
	if err == redis.Nil || storedVal != req.VerificationToken {
		return errors.New(string(i18n.KeyResetTokenExpired))
	}
	user, err := u.Repo.FindUserByEmail(ctx, email)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	if err := u.Repo.UpdateUser(ctx, user); err != nil {
		return err
	}
	u.RDB.Del(ctx, verifyTokenKey)
	return nil
}
