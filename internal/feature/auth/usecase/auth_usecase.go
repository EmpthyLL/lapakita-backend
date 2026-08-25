package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/repository"
	"lapakita-backend/pkg/i18n"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/mailer"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type AuthUsecase struct {
	repo       repository.AuthRepository
	jwtService jwt.JWTService
	mailer     mailer.Mailer
}

func NewAuthUsecase(
	repo repository.AuthRepository,
	jwtService jwt.JWTService,
	mailer mailer.Mailer,
) *AuthUsecase {
	return &AuthUsecase{
		repo:       repo,
		jwtService: jwtService,
		mailer:     mailer,
	}
}

func generateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	return fmt.Sprintf("%06d", n.Int64()+100000)
}

func (u *AuthUsecase) helperBuildAuthResponse(user *entity.User) (*dto.AuthResponseData, error) {
	primaryPhone := user.PhoneNumbers.GetPrimaryNumber()

	var avatarPtr *string
	if user.DefaultAvatarURL != nil && *user.DefaultAvatarURL != "" {
		avatarPtr = user.DefaultAvatarURL
	}

	var subExpiresAtPtr *string
	subExpiresAtStr := ""
	if user.SubscriptionExpiresAt != nil {
		formatted := user.SubscriptionExpiresAt.Format(time.RFC3339)
		subExpiresAtPtr = &formatted
		subExpiresAtStr = formatted
	}

	avatarStr := ""
	if avatarPtr != nil {
		avatarStr = *avatarPtr
	}

	claims := jwt.JWTCustomClaims{
		UserID:                user.ID.String(),
		Name:                  user.Name,
		Email:                 user.Email,
		Phone:                 primaryPhone,
		AvatarURL:             avatarStr,
		ActiveRole:            user.ActiveRole,
		SubscriptionPlan:      user.SubscriptionPlan,
		SubscriptionExpiresAt: subExpiresAtStr,
	}

	tokenPair, err := u.jwtService.GenerateTokenPair(claims)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponseData{
		User: dto.UserPayload{
			ID:                    user.ID.String(),
			Name:                  user.Name,
			Email:                 user.Email,
			Phone:                 primaryPhone,
			AvatarURL:             avatarPtr,
			ActiveRole:            user.ActiveRole,
			SubscriptionPlan:      user.SubscriptionPlan,
			SubscriptionExpiresAt: subExpiresAtPtr,
			Token:                 tokenPair.AccessToken,
		},
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

// -----------------------------------------------------------------------------
// 1. REGISTER
// -----------------------------------------------------------------------------
func (u *AuthUsecase) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponseData, error) {
	existingUser, _ := u.repo.FindUserByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New(string(i18n.KeyEmailAlreadyRegistered))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		ActiveRole:   "tenant",
		PhoneNumbers: entity.PhoneNumbers{
			{
				Number:    req.Phone,
				IsPrimary: true,
				Roles:     []string{},
			},
		},
	}

	if err := u.repo.CreateUser(ctx, newUser); err != nil {
		return nil, err
	}

	return u.helperBuildAuthResponse(newUser)
}

// -----------------------------------------------------------------------------
// 2. LOGIN
// -----------------------------------------------------------------------------
func (u *AuthUsecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponseData, error) {
	user, err := u.repo.FindUserByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New(string(i18n.KeyWrongPassword))
	}

	return u.helperBuildAuthResponse(user)
}

// -----------------------------------------------------------------------------
// 3. GOOGLE AUTH & COMPLETE PROFILE
// -----------------------------------------------------------------------------
func (u *AuthUsecase) GoogleAuth(ctx context.Context, req dto.GoogleAuthRequest) (*dto.AuthResponseData, *dto.GoogleSetupPresetResponse, error) {
	payload, err := idtoken.Validate(ctx, req.IDToken, "")
	if err != nil {
		return nil, nil, errors.New(string(i18n.KeyGoogleAuthFailed))
	}

	email := fmt.Sprintf("%v", payload.Claims["email"])
	name := fmt.Sprintf("%v", payload.Claims["name"])
	picture := fmt.Sprintf("%v", payload.Claims["picture"])

	user, _ := u.repo.FindUserByEmail(ctx, email)

	if user == nil {
		newUser := &entity.User{
			Name:             name,
			Email:            email,
			PasswordHash:     "",
			DefaultAvatarURL: &picture,
			ActiveRole:       "tenant",
			PhoneNumbers:     entity.PhoneNumbers{},
		}
		if err := u.repo.CreateUser(ctx, newUser); err != nil {
			return nil, nil, err
		}

		setupData := dto.GoogleTokenPayload{Email: email, Name: name, AvatarURL: picture}
		jsonBytes, _ := json.Marshal(setupData)
		setupToken := base64.StdEncoding.EncodeToString(jsonBytes)

		return nil, &dto.GoogleSetupPresetResponse{
			Email:              email,
			Name:               name,
			AvatarURL:          picture,
			SetupToken:         setupToken,
			IsProfileCompleted: false,
		}, nil
	}

	if user.PasswordHash == "" || len(user.PhoneNumbers) == 0 {
		setupData := dto.GoogleTokenPayload{Email: email, Name: name, AvatarURL: picture}
		jsonBytes, _ := json.Marshal(setupData)
		setupToken := base64.StdEncoding.EncodeToString(jsonBytes)

		return nil, &dto.GoogleSetupPresetResponse{
			Email:              email,
			Name:               name,
			AvatarURL:          picture,
			SetupToken:         setupToken,
			IsProfileCompleted: false,
		}, nil
	}

	authRes, err := u.helperBuildAuthResponse(user)
	return authRes, nil, err
}

func (u *AuthUsecase) CompleteProfile(ctx context.Context, req dto.CompleteProfileRequest) (*dto.AuthResponseData, error) {
	rawBytes, err := base64.StdEncoding.DecodeString(req.SetupToken)
	if err != nil {
		return nil, errors.New(string(i18n.KeySetupTokenExpired))
	}

	var googleData dto.GoogleTokenPayload
	if err := json.Unmarshal(rawBytes, &googleData); err != nil {
		return nil, errors.New(string(i18n.KeySetupTokenExpired))
	}

	user, err := u.repo.FindUserByEmail(ctx, googleData.Email)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	avatar := req.AvatarURL
	if avatar == "" {
		avatar = googleData.AvatarURL
	}

	user.Name = req.Name
	user.PasswordHash = string(hashedPassword)
	user.DefaultAvatarURL = &avatar
	user.PhoneNumbers = entity.PhoneNumbers{
		{
			Number:    req.Phone,
			IsPrimary: true,
			Roles:     []string{},
		},
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return u.helperBuildAuthResponse(user)
}

// -----------------------------------------------------------------------------
// 4. OTP SYSTEM
// -----------------------------------------------------------------------------
func (u *AuthUsecase) SendOTP(ctx context.Context, req dto.SendOTPRequest) (string, error) {
	otpCode := generateOTP()

	if err := u.repo.SetOTP(ctx, req.Mode, req.Email, otpCode, 5*time.Minute); err != nil {
		return "", err
	}

	go func() {
		_ = u.mailer.SendOTPEmail(req.Email, otpCode, req.Mode)
	}()

	stateData := dto.OTPStatePayload{Email: req.Email, Mode: req.Mode}
	jsonBytes, _ := json.Marshal(stateData)
	encodedState := base64.StdEncoding.EncodeToString(jsonBytes)

	return encodedState, nil
}

func (u *AuthUsecase) VerifyOTP(ctx context.Context, req dto.VerifyOTPRequest) (string, error) {
	rawBytes, err := base64.StdEncoding.DecodeString(req.StatePayload)
	if err != nil {
		return "", errors.New(string(i18n.KeyOTPExpired))
	}

	var state dto.OTPStatePayload
	if err := json.Unmarshal(rawBytes, &state); err != nil {
		return "", errors.New(string(i18n.KeyOTPExpired))
	}

	storedOTP, err := u.repo.GetOTP(ctx, state.Mode, state.Email)
	if err == redis.Nil || storedOTP != req.OTPCode {
		return "", errors.New(string(i18n.KeyOTPExpired))
	}

	_ = u.repo.DeleteOTP(ctx, state.Mode, state.Email)

	tokenVal := fmt.Sprintf("vtok_%d", time.Now().UnixNano())
	if err := u.repo.SetVerificationToken(ctx, state.Mode, state.Email, tokenVal, 15*time.Minute); err != nil {
		return "", err
	}

	return tokenVal, nil
}

// -----------------------------------------------------------------------------
// 5. RESET PASSWORD & REFRESH TOKEN
// -----------------------------------------------------------------------------
func (u *AuthUsecase) ResetPassword(ctx context.Context, email string, req dto.ResetPasswordRequest) error {
	storedVal, err := u.repo.GetVerificationToken(ctx, "reset_password", email)
	if err == redis.Nil || storedVal != req.VerificationToken {
		return errors.New(string(i18n.KeyResetTokenExpired))
	}

	user, err := u.repo.FindUserByEmail(ctx, email)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return err
	}

	_ = u.repo.DeleteVerificationToken(ctx, "reset_password", email)
	return nil
}

func (u *AuthUsecase) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (*dto.AuthResponseData, error) {
	parsedToken, err := u.jwtService.ValidateToken(req.RefreshToken)
	if err != nil || !parsedToken.Valid {
		return nil, errors.New(string(i18n.KeyRefreshTokenExpired))
	}

	claims, ok := parsedToken.Claims.(*jwt.JWTCustomClaims)
	if !ok || claims.TokenType != "refresh" {
		return nil, errors.New(string(i18n.KeyRefreshTokenExpired))
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, errors.New(string(i18n.KeyRefreshTokenExpired))
	}

	user, err := u.repo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}

	return u.helperBuildAuthResponse(user)
}
