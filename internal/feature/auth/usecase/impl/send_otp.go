package impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/pkg/i18n"
)

func SendOTP(ctx context.Context, u *usecase.AuthUsecase, req dto.SendOTPRequest) error {
	switch req.Mode {
	case "register":
		pendingKey := fmt.Sprintf("pending_reg:%s", req.Email)
		exists, err := u.RDB.Exists(ctx, pendingKey).Result()
		if err != nil || exists == 0 {
			return errors.New(string(i18n.KeyUserNotFound))
		}
	case "reset_password":
		user, err := u.Repo.FindUserByEmail(ctx, req.Email)
		if err != nil || user == nil {
			return errors.New(string(i18n.KeyUserNotFound))
		}
	}
	otpCode := usecase.GenerateOTP()
	redisKey := fmt.Sprintf("otp:%s:%s", req.Mode, req.Email)
	if err := u.RDB.Set(ctx, redisKey, otpCode, 5*time.Minute).Err(); err != nil {
		return err
	}
	go func(email, code, mode string) { _ = u.Mailer.SendOTPEmail(email, code, mode) }(req.Email, otpCode, req.Mode)
	return nil
}
