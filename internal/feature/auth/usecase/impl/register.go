package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/pkg/i18n"

	"golang.org/x/crypto/bcrypt"
)

func Register(ctx context.Context, u *usecase.AuthUsecase, req dto.RegisterRequest) error {
	existingUser, _ := u.Repo.FindUserByEmail(ctx, req.Email)
	if existingUser != nil {
		return errors.New(string(i18n.KeyEmailAlreadyRegistered))
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	pendingUserKey := fmt.Sprintf("pending_reg:%s", req.Email)
	pendingData, _ := json.Marshal(map[string]string{"name": req.Name, "email": req.Email, "password_hash": string(hashedPassword), "dial_code": req.DialCode, "phone": req.Phone})
	if err := u.RDB.Set(ctx, pendingUserKey, pendingData, 10*time.Minute).Err(); err != nil {
		return err
	}
	otpCode := usecase.GenerateOTP()
	otpKey := fmt.Sprintf("otp:register:%s", req.Email)
	if err := u.RDB.Set(ctx, otpKey, otpCode, 5*time.Minute).Err(); err != nil {
		return err
	}
	go func(email, code string) { _ = u.Mailer.SendOTPEmail(email, code, "register") }(req.Email, otpCode)
	return nil
}
