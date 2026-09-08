package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/repository"
	"lapakita-backend/pkg/i18n"
	"lapakita-backend/pkg/storage"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo     *repository.UserRepository
	imagekit *storage.ImageKitService
}

func NewUserUsecase(repo *repository.UserRepository, imagekit *storage.ImageKitService) *UserUsecase {
	return &UserUsecase{
		repo:     repo,
		imagekit: imagekit,
	}
}

// 1. General Profile
func (u *UserUsecase) GetGeneralProfile(ctx context.Context, userID uuid.UUID) (dto.GetGeneralProfileResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetGeneralProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	avatar := ""
	if user.DefaultAvatarURL != nil {
		avatar = *user.DefaultAvatarURL
	}

	return dto.GetGeneralProfileResponse{
		ID:               user.ID.String(),
		Name:             user.Name,
		Email:            user.Email,
		DefaultAvatarURL: avatar,
		PrimaryPhone:     user.PhoneNumbers.GetPrimaryNumber(),
	}, nil
}

func (u *UserUsecase) UpdateGeneralProfile(ctx context.Context, userID uuid.UUID, req dto.UpdateGeneralProfileRequest) (dto.GetGeneralProfileResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetGeneralProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	// 1. Update nama & avatar
	user.Name = req.Name
	if req.DefaultAvatarURL != nil {
		user.DefaultAvatarURL = req.DefaultAvatarURL
	}

	// 2. Jika phone_number dikirim di payload, set/tambahkan dan jadikan primary
	if req.PhoneNumber != nil && *req.PhoneNumber != "" {
		phone := *req.PhoneNumber
		found := false

		// Cek apakah nomor sudah ada di daftar
		for i := range user.PhoneNumbers {
			if user.PhoneNumbers[i].Number == phone {
				user.PhoneNumbers[i].IsPrimary = true
				found = true
			} else {
				user.PhoneNumbers[i].IsPrimary = false
			}
		}

		// Jika nomor baru, tambahkan ke list dan set sebagai primary
		if !found {
			// Unset primary nomor lama
			for i := range user.PhoneNumbers {
				user.PhoneNumbers[i].IsPrimary = false
			}

			user.PhoneNumbers = append(user.PhoneNumbers, entity.PhoneNumberItem{
				Number:    phone,
				IsPrimary: true,
				Roles:     []string{"primary_contact"},
			})
		}
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return dto.GetGeneralProfileResponse{}, err
	}

	return u.GetGeneralProfile(ctx, userID)
}

// 2. Phone Numbers
func (u *UserUsecase) GetPhoneNumbers(ctx context.Context, userID uuid.UUID) (dto.GetPhoneNumbersResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetPhoneNumbersResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	items := make([]dto.PhoneNumberItem, 0, len(user.PhoneNumbers))
	for _, p := range user.PhoneNumbers {
		items = append(items, dto.PhoneNumberItem{
			Number:    p.Number,
			IsPrimary: p.IsPrimary,
			Roles:     p.Roles,
		})
	}

	return dto.GetPhoneNumbersResponse{PhoneNumbers: items}, nil
}

func (u *UserUsecase) AddPhoneNumber(ctx context.Context, userID uuid.UUID, req dto.AddPhoneNumberRequest) (dto.GetPhoneNumbersResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetPhoneNumbersResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	for _, p := range user.PhoneNumbers {
		if p.Number == req.Number {
			return dto.GetPhoneNumbersResponse{}, errors.New(string(i18n.KeyUserPhoneDuplicate))
		}
	}

	newItem := entity.PhoneNumberItem{
		Number:    req.Number,
		IsPrimary: req.IsPrimary,
		Roles:     req.Roles,
	}

	if req.IsPrimary || len(user.PhoneNumbers) == 0 {
		newItem.IsPrimary = true
		for i := range user.PhoneNumbers {
			user.PhoneNumbers[i].IsPrimary = false
		}
	}

	user.PhoneNumbers = append(user.PhoneNumbers, newItem)
	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return dto.GetPhoneNumbersResponse{}, err
	}

	return u.GetPhoneNumbers(ctx, userID)
}

func (u *UserUsecase) UpdatePhoneNumber(ctx context.Context, userID uuid.UUID, index int, req dto.UpdatePhoneNumberRequest) (dto.GetPhoneNumbersResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetPhoneNumbersResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	if index < 0 || index >= len(user.PhoneNumbers) {
		return dto.GetPhoneNumbersResponse{}, errors.New(string(i18n.KeyUserPhoneIndexInvalid))
	}

	for i, p := range user.PhoneNumbers {
		if i != index && p.Number == req.Number {
			return dto.GetPhoneNumbersResponse{}, errors.New(string(i18n.KeyUserPhoneDuplicate))
		}
	}

	user.PhoneNumbers[index].Number = req.Number
	user.PhoneNumbers[index].Roles = req.Roles

	if req.IsPrimary {
		for i := range user.PhoneNumbers {
			user.PhoneNumbers[i].IsPrimary = (i == index)
		}
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return dto.GetPhoneNumbersResponse{}, err
	}

	return u.GetPhoneNumbers(ctx, userID)
}

func (u *UserUsecase) DeletePhoneNumber(ctx context.Context, userID uuid.UUID, index int) error {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}

	if index < 0 || index >= len(user.PhoneNumbers) {
		return errors.New(string(i18n.KeyUserPhoneIndexInvalid))
	}

	if user.PhoneNumbers[index].IsPrimary {
		return errors.New(string(i18n.KeyUserPhoneCannotDeletePrimary))
	}

	user.PhoneNumbers = append(user.PhoneNumbers[:index], user.PhoneNumbers[index+1:]...)
	return u.repo.UpdateUser(ctx, user)
}

// 3. Security / Password
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

// 4. Persona Profile
func (u *UserUsecase) GetPersonaProfile(ctx context.Context, userID uuid.UUID, role string) (dto.PersonaProfileResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.PersonaProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	profile, exists := user.RoleProfiles[role]
	if !exists {
		defaultAvatar := ""
		if user.DefaultAvatarURL != nil {
			defaultAvatar = *user.DefaultAvatarURL
		}
		return dto.PersonaProfileResponse{
			Role:        role,
			DisplayName: user.Name,
			AvatarURL:   defaultAvatar,
		}, nil
	}

	return dto.PersonaProfileResponse{
		Role:        role,
		DisplayName: profile.DisplayName,
		AvatarURL:   profile.AvatarURL,
	}, nil
}

func (u *UserUsecase) UpdatePersonaProfile(ctx context.Context, userID uuid.UUID, role string, req dto.UpdatePersonaRequest) (dto.PersonaProfileResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.PersonaProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	if user.RoleProfiles == nil {
		user.RoleProfiles = entity.RoleProfiles{}
	}

	user.RoleProfiles[role] = entity.RoleProfileItem{
		DisplayName: req.DisplayName,
		AvatarURL:   req.AvatarURL,
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return dto.PersonaProfileResponse{}, err
	}

	return u.GetPersonaProfile(ctx, userID, role)
}

// 5. Document Upload & Watermarking
func (u *UserUsecase) UploadDocument(ctx context.Context, userID uuid.UUID, req dto.UploadDocumentRequest) (dto.GetDocumentResponse, error) {
	// 1. Cek Unik NIK
	existingIdentity, err := u.repo.FindIdentityByNIK(ctx, req.NIK)
	if err != nil {
		return dto.GetDocumentResponse{}, err
	}
	if existingIdentity != nil && existingIdentity.UserID != userID {
		return dto.GetDocumentResponse{}, errors.New(string(i18n.KeyUserDocumentNIKExists))
	}

	// 2. Clean Base64 String
	rawBase64 := req.KTPPhoto
	if idx := strings.Index(rawBase64, ","); idx != -1 {
		rawBase64 = rawBase64[idx+1:]
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(rawBase64)
	if err != nil {
		return dto.GetDocumentResponse{}, errors.New(string(i18n.KeyUserDocumentFileInvalid))
	}

	// 3. Set Purpose otomatis dari BE (misal: Stall/User Verification)
	purpose := storage.PurposeStallVerification
	watermarkedBytes, err := storage.ApplyWatermarkFromBytes(decodedBytes, purpose)
	if err != nil {
		return dto.GetDocumentResponse{}, errors.New(string(i18n.KeyUserDocumentWatermarkFailed))
	}

	// 4. Upload ke ImageKit via Base64
	watermarkedBase64 := base64.StdEncoding.EncodeToString(watermarkedBytes)
	fileName := fmt.Sprintf("ktp_%s.png", userID.String())

	uploadedURL, err := u.imagekit.UploadFromURL(ctx, watermarkedBase64, fileName, "/users/identity_documents")
	if err != nil {
		return dto.GetDocumentResponse{}, errors.New(string(i18n.KeyUserDocumentFailedToUpload))
	}

	// 5. Simpan / Upsert ke Database
	domicile := req.DomicileCity
	identity := &entity.UserIdentityProfile{
		UserID:       userID,
		FullNameKTP:  req.FullNameKTP,
		NIK:          req.NIK,
		KTPPhotoURL:  uploadedURL,
		DomicileCity: &domicile,
	}

	if err := u.repo.UpsertIdentityProfile(ctx, identity); err != nil {
		return dto.GetDocumentResponse{}, err
	}

	return dto.GetDocumentResponse{
		FullNameKTP:  identity.FullNameKTP,
		NIK:          identity.NIK,
		KTPPhotoURL:  identity.KTPPhotoURL,
		DomicileCity: domicile,
	}, nil
}

func (u *UserUsecase) DeleteDocument(ctx context.Context, userID uuid.UUID) error {
	return u.repo.DeleteIdentityProfile(ctx, userID)
}
