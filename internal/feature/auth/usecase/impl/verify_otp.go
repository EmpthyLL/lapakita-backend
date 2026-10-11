package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (u *AuthUsecase) VerifyOTP(ctx context.Context, req dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error) {
	redisKey := fmt.Sprintf("otp:%s:%s", req.Mode, req.Email)
	storedOTP, err := u.rdb.Get(ctx, redisKey).Result()
	if err == redis.Nil || storedOTP != req.OTPCode {
		return nil, errors.New(string(i18n.KeyOTPExpired))
	}
	if req.Mode == "register" {
		pendingKey := fmt.Sprintf("pending_reg:%s", req.Email)
		pendingStr, err := u.rdb.Get(ctx, pendingKey).Result()
		if err == redis.Nil {
			return nil, errors.New(string(i18n.KeyOTPExpired))
		}
		var regData map[string]string
		if err := json.Unmarshal([]byte(pendingStr), &regData); err != nil {
			return nil, errors.New(string(i18n.KeyOTPExpired))
		}
		newUser := &entity.User{
			Name: regData["name"], Email: regData["email"], PasswordHash: regData["password_hash"], ActiveRole: "tenant", SubscriptionPlan: "free",
			PhoneNumbers: entity.PhoneNumbers{{DialCode: regData["dial_code"], Number: regData["phone"], IsPrimary: true, Roles: []string{}, Label: "Primary Number"}}, RoleProfiles: entity.RoleProfiles{},
		}
		if err := u.repo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}
		u.rdb.Del(ctx, redisKey)
		u.rdb.Del(ctx, pendingKey)
		authRes, err := usecase.BuildAuthResponse(u.jwtService, newUser)
		if err != nil {
			return nil, err
		}
		return &dto.VerifyOTPResponse{AuthData: authRes}, nil
	}
	u.rdb.Del(ctx, redisKey)
	tokenVal := uuid.New().String()
	verifyTokenKey := fmt.Sprintf("verified:%s:%s", req.Mode, req.Email)
	if err := u.rdb.Set(ctx, verifyTokenKey, tokenVal, 15*time.Minute).Err(); err != nil {
		return nil, err
	}
	return &dto.VerifyOTPResponse{VerificationToken: tokenVal}, nil
}
