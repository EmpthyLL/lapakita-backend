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
	imagekit   *storage.ImageKitService // Added ImageKit Dependency
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

func (u *AuthUsecase) helperBuildAuthResponse(user *entity.User) (*dto.AuthResponseData, error) {
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

	userPayload := dto.UserPayload{
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
	// 1. Cek ketersediaan email di PostgreSQL
	existingUser, _ := u.repo.FindUserByEmail(ctx, req.Email)
	if existingUser != nil {
		return errors.New(string(i18n.KeyEmailAlreadyRegistered))
	}

	// 2. Hash Password & simpan pending registration ke Redis (TTL 10 menit)
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

	// 3. Send OTP (Simpan dengan format key: otp:register:<email>)
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

	// Helper closure untuk upload ke ImageKit
	uploadAvatar := func(rawURL string) string {
		if rawURL == "" {
			return ""
		}
		fileName := fmt.Sprintf("avatar_%s.jpg", uuid.New().String())
		ikURL, err := u.imagekit.UploadFromURL(ctx, rawURL, fileName, "/avatars")
		if err != nil {
			fmt.Printf("[ImageKit Upload Error]: %v\n", err) // Log error jika upload gagal
			return rawURL                                    // Fallback ke URL asli jika gagal
		}
		return ikURL
	}

	if user == nil {
		var avatarPtr *string
		if picture != "" {
			ikURL := uploadAvatar(picture)
			avatarPtr = &ikURL
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
		// Jika user lama masih memiliki URL avatar Google asli, update ke ImageKit
		if user.DefaultAvatarURL != nil && strings.Contains(*user.DefaultAvatarURL, "googleusercontent.com") {
			ikURL := uploadAvatar(picture)
			if ikURL != *user.DefaultAvatarURL {
				user.DefaultAvatarURL = &ikURL
				_ = u.repo.UpdateUser(ctx, user)
			}
		}
	}

	return u.helperBuildAuthResponse(user)
}

func (u *AuthUsecase) CompleteProfile(ctx context.Context, userID uuid.UUID, req dto.CompleteProfileRequest) (*dto.AuthResponseData, error) {
	// 1. Cari user di DB berdasarkan User ID dari JWT Claims Session
	user, err := u.repo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}

	// 2. Update profil user
	user.Name = req.Name
	if req.AvatarURL != "" {
		user.DefaultAvatarURL = &req.AvatarURL
	}

	// Update / Set primary phone number
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

	// 3. Re-build Auth Response dengan data profil yang terbaru
	return u.helperBuildAuthResponse(user)
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
	// 1. Verifikasi Kode OTP dari Redis (key: otp:<mode>:<email>)
	redisKey := fmt.Sprintf("otp:%s:%s", req.Mode, req.Email)
	storedOTP, err := u.rdb.Get(ctx, redisKey).Result()
	if err == redis.Nil || storedOTP != req.OTPCode {
		return nil, errors.New(string(i18n.KeyOTPExpired))
	}
	// 2. Jika Mode REGISTER -> Ambil pending data dari Redis & simpan user ke DB
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

		// Hapus Redis key
		u.rdb.Del(ctx, redisKey)
		u.rdb.Del(ctx, pendingKey)

		// Generate token JWT & Auth Response
		authRes, err := u.helperBuildAuthResponse(newUser)
		if err != nil {
			return nil, err
		}

		// Pastikan mengembalikan &dto.VerifyOTPResponse dengan AuthData yang terisi
		return &dto.VerifyOTPResponse{
			AuthData: authRes,
		}, nil
	}

	// 3. Jika Mode RESET PASSWORD -> Hasilkan verification token UUID untuk reset password
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
