package usecase

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/repository"
	"lapakita-backend/pkg/i18n"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/mailer"
	"lapakita-backend/pkg/storage"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type AuthUsecase struct {
	repo       *repository.AuthRepository
	rdb        *redis.Client
	jwtService *jwt.JWTService
	mailer     *mailer.Mailer
	imagekit   *storage.ImageKitService
}

func NewAuthUsecase(
	repo *repository.AuthRepository,
	rdb *redis.Client,
	jwtService *jwt.JWTService,
	mailer *mailer.Mailer,
	imagekit *storage.ImageKitService,
) *AuthUsecase {
	return &AuthUsecase{
		repo:       repo,
		rdb:        rdb,
		jwtService: jwtService,
		mailer:     mailer,
		imagekit:   imagekit,
	}
}

func generateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	return fmt.Sprintf("%06d", n.Int64()+100000)
}

// helperBuildUserPayload khusus menyusun UserPayload tanpa melakukan pembuatan Token JWT baru
func (u *AuthUsecase) helperBuildUserPayload(user *entity.User) dto.UserPayload {
	defaultPhone := user.PhoneNumbers.GetPrimaryNumber()

	var defaultAvatarPtr *string
	defaultAvatarStr := ""
	if user.DefaultAvatarURL != nil && *user.DefaultAvatarURL != "" {
		defaultAvatarPtr = user.DefaultAvatarURL
		defaultAvatarStr = *user.DefaultAvatarURL
	}

	activePlan := user.SubscriptionPlan
	var subExpiresAtPtr *string

	if user.SubscriptionExpiresAt != nil {
		if time.Now().After(*user.SubscriptionExpiresAt) {
			activePlan = "free"
		} else {
			formatted := user.SubscriptionExpiresAt.Format(time.RFC3339)
			subExpiresAtPtr = &formatted
		}
	}

	phonePayloads := make([]dto.PhonePayload, 0)
	roleToPhoneMap := make(map[string]string)

	for _, p := range user.PhoneNumbers {
		phonePayloads = append(phonePayloads, dto.PhonePayload{
			Number:    p.Number,
			IsPrimary: p.IsPrimary,
			Roles:     p.Roles,
		})

		for _, role := range p.Roles {
			if _, exists := roleToPhoneMap[role]; !exists {
				roleToPhoneMap[role] = p.Number
			}
		}
	}

	personas := make(map[string]dto.PersonaDetail)
	for roleKey, profile := range user.RoleProfiles {
		phoneVal := roleToPhoneMap[roleKey]
		if phoneVal == "" {
			phoneVal = defaultPhone
		}

		avatarVal := profile.AvatarURL
		if avatarVal == "" {
			avatarVal = defaultAvatarStr
		}

		nameVal := profile.DisplayName
		if nameVal == "" {
			nameVal = user.Name
		}

		personas[roleKey] = dto.PersonaDetail{
			DisplayName: nameVal,
			AvatarURL:   avatarVal,
			Phone:       phoneVal,
		}
	}

	return dto.UserPayload{
		ID:                    user.ID.String(),
		DefaultName:           user.Name,
		DefaultAvatarURL:      defaultAvatarPtr,
		DefaultPhone:          defaultPhone,
		Email:                 user.Email,
		IsPasswordSet:         user.PasswordHash != "",
		ActiveRole:            user.ActiveRole,
		SubscriptionPlan:      activePlan,
		SubscriptionExpiresAt: subExpiresAtPtr,
		PhoneNumbers:          phonePayloads,
		Personas:              personas,
	}
}

// helperBuildAuthResponse digunakan khusus saat INIT SESSION (Login/Register/GoogleAuth/RefreshToken)
func (u *AuthUsecase) helperBuildAuthResponse(user *entity.User) (*dto.AuthResponseData, error) {
	userPayload := u.helperBuildUserPayload(user)

	claims := jwt.JWTCustomClaims{
		UserPayload: userPayload,
	}

	tokenPair, err := u.jwtService.GenerateTokenPair(claims)
	if err != nil {
		return nil, err
	}

	userPayload.Token = tokenPair.AccessToken

	return &dto.AuthResponseData{
		User:         userPayload,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

// -----------------------------------------------------------------------------
// 1. REGISTER & LOGIN
// -----------------------------------------------------------------------------
func (u *AuthUsecase) Register(ctx context.Context, req dto.RegisterRequest) error {
	existingUser, _ := u.repo.FindUserByEmail(ctx, req.Email)
	if existingUser != nil {
		return errors.New(string(i18n.KeyEmailAlreadyRegistered))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pendingUserKey := fmt.Sprintf("pending_reg:%s", req.Email)
	pendingData, _ := json.Marshal(map[string]string{
		"name":          req.Name,
		"email":         req.Email,
		"password_hash": string(hashedPassword),
		"phone":         req.Phone,
	})

	if err := u.rdb.Set(ctx, pendingUserKey, pendingData, 10*time.Minute).Err(); err != nil {
		return err
	}

	otpCode := generateOTP()
	otpKey := fmt.Sprintf("otp:register:%s", req.Email)
	if err := u.rdb.Set(ctx, otpKey, otpCode, 5*time.Minute).Err(); err != nil {
		return err
	}

	go func(email, code string) {
		_ = u.mailer.SendOTPEmail(email, code, "register")
	}(req.Email, otpCode)

	return nil
}

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
// 2. GOOGLE AUTH & COMPLETE PROFILE
// -----------------------------------------------------------------------------
func (u *AuthUsecase) GoogleAuth(ctx context.Context, req dto.GoogleAuthRequest) (*dto.AuthResponseData, error) {
	payload, err := idtoken.Validate(ctx, req.IDToken, "")
	if err != nil {
		return nil, errors.New(string(i18n.KeyGoogleAuthFailed))
	}

	email := fmt.Sprintf("%v", payload.Claims["email"])
	name := fmt.Sprintf("%v", payload.Claims["name"])
	picture := fmt.Sprintf("%v", payload.Claims["picture"])

	user, _ := u.repo.FindUserByEmail(ctx, email)

	// Helper closure untuk upload ke ImageKit (mendukung URL HTTP & Base64)
	uploadAvatar := func(rawSource string) string {
		cleanSource := strings.TrimSpace(rawSource)
		if cleanSource == "" {
			return ""
		}

		fileName := fmt.Sprintf("avatar_%s.jpg", uuid.New().String())
		ikURL, err := u.imagekit.UploadFromURL(ctx, cleanSource, fileName, "/avatars")
		if err != nil {
			fmt.Printf("[ImageKit Upload Error]: %v\n", err)
			return cleanSource // Fallback ke string asli jika gagal
		}
		return ikURL
	}

	if user == nil {
		var avatarPtr *string
		if picture != "" {
			ikURL := uploadAvatar(picture)
			if ikURL != "" {
				avatarPtr = &ikURL
			}
		}

		newUser := &entity.User{
			Name:             name,
			Email:            email,
			PasswordHash:     "",
			DefaultAvatarURL: avatarPtr,
			ActiveRole:       "tenant",
			SubscriptionPlan: "free",
			PhoneNumbers:     entity.PhoneNumbers{},
			RoleProfiles:     entity.RoleProfiles{},
		}

		if err := u.repo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}
		user = newUser
	} else {
		needsUpload := user.DefaultAvatarURL == nil ||
			*user.DefaultAvatarURL == "" ||
			strings.Contains(*user.DefaultAvatarURL, "googleusercontent.com") ||
			strings.HasPrefix(*user.DefaultAvatarURL, "data:image") ||
			(!strings.HasPrefix(*user.DefaultAvatarURL, "http://") && !strings.HasPrefix(*user.DefaultAvatarURL, "https://"))

		if needsUpload && picture != "" {
			ikURL := uploadAvatar(picture)
			if ikURL != "" && (user.DefaultAvatarURL == nil || ikURL != *user.DefaultAvatarURL) {
				user.DefaultAvatarURL = &ikURL
				_ = u.repo.UpdateUser(ctx, user)
			}
		}
	}

	// Google Auth memicu inisialisasi session -> return AuthResponseData (dengan Token)
	return u.helperBuildAuthResponse(user)
}

func (u *AuthUsecase) CompleteProfile(ctx context.Context, userID uuid.UUID, req dto.CompleteProfileRequest) (*dto.UserPayload, error) {
	user, err := u.repo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}

	user.Name = req.Name

	// Jika avatar diisi (bisa berupa URL atau Base64), upload ke ImageKit
	if req.AvatarURL != "" {
		avatarSource := strings.TrimSpace(req.AvatarURL)

		// Upload ke ImageKit jika bertipe Base64 atau URL luar
		if !strings.Contains(avatarSource, "ik.imagekit.io") {
			fileName := fmt.Sprintf("avatar_%s.jpg", uuid.New().String())
			ikURL, err := u.imagekit.UploadFromURL(ctx, avatarSource, fileName, "/avatars")
			if err == nil && ikURL != "" {
				avatarSource = ikURL
			}
		}
		user.DefaultAvatarURL = &avatarSource
	}

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

	updatedPayload := u.helperBuildUserPayload(user)
	return &updatedPayload, nil
}

// -----------------------------------------------------------------------------
// 3. OTP SYSTEM
// -----------------------------------------------------------------------------
func (u *AuthUsecase) SendOTP(ctx context.Context, req dto.SendOTPRequest) error {
	switch req.Mode {
	case "register":
		pendingKey := fmt.Sprintf("pending_reg:%s", req.Email)
		exists, err := u.rdb.Exists(ctx, pendingKey).Result()
		if err != nil || exists == 0 {
			return errors.New(string(i18n.KeyUserNotFound))
		}
	case "reset_password":
		user, err := u.repo.FindUserByEmail(ctx, req.Email)
		if err != nil || user == nil {
			return errors.New(string(i18n.KeyUserNotFound))
		}
	}

	otpCode := generateOTP()
	redisKey := fmt.Sprintf("otp:%s:%s", req.Mode, req.Email)

	if err := u.rdb.Set(ctx, redisKey, otpCode, 5*time.Minute).Err(); err != nil {
		return err
	}

	go func(email, code, mode string) {
		_ = u.mailer.SendOTPEmail(email, code, mode)
	}(req.Email, otpCode, req.Mode)

	return nil
}

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
			Name:             regData["name"],
			Email:            regData["email"],
			PasswordHash:     regData["password_hash"],
			ActiveRole:       "tenant",
			SubscriptionPlan: "free",
			PhoneNumbers: entity.PhoneNumbers{
				{
					Number:    regData["phone"],
					IsPrimary: true,
					Roles:     []string{},
				},
			},
			RoleProfiles: entity.RoleProfiles{},
		}

		if err := u.repo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}

		u.rdb.Del(ctx, redisKey)
		u.rdb.Del(ctx, pendingKey)

		// Pendaftaran via OTP memicu Login Pertama -> Return AuthResponseData (dengan Token)
		authRes, err := u.helperBuildAuthResponse(newUser)
		if err != nil {
			return nil, err
		}

		return &dto.VerifyOTPResponse{
			AuthData: authRes,
		}, nil
	}

	u.rdb.Del(ctx, redisKey)

	tokenVal := uuid.New().String()
	verifyTokenKey := fmt.Sprintf("verified:%s:%s", req.Mode, req.Email)
	if err := u.rdb.Set(ctx, verifyTokenKey, tokenVal, 15*time.Minute).Err(); err != nil {
		return nil, err
	}

	return &dto.VerifyOTPResponse{
		VerificationToken: tokenVal,
	}, nil
}

// -----------------------------------------------------------------------------
// 4. RESET PASSWORD & REFRESH TOKEN
// -----------------------------------------------------------------------------
func (u *AuthUsecase) ResetPassword(ctx context.Context, email string, req dto.ResetPasswordRequest) error {
	verifyTokenKey := fmt.Sprintf("verified:reset_password:%s", email)
	storedVal, err := u.rdb.Get(ctx, verifyTokenKey).Result()
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

	u.rdb.Del(ctx, verifyTokenKey)
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

	userID, err := uuid.Parse(claims.UserPayload.ID)
	if err != nil {
		return nil, errors.New(string(i18n.KeyRefreshTokenExpired))
	}

	user, err := u.repo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}

	return u.helperBuildAuthResponse(user)
}
