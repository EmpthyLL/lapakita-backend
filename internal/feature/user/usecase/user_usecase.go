package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/repository"
	"lapakita-backend/pkg/api"
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
		ActiveRole:       user.ActiveRole,
		PrimaryPhone:     user.PhoneNumbers.GetPrimaryNumber(),
	}, nil
}

func (u *UserUsecase) UpdateGeneralProfile(ctx context.Context, userID uuid.UUID, req dto.UpdateGeneralProfileRequest) (dto.GetGeneralProfileResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetGeneralProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}

	user.Name = req.Name

	if req.DefaultAvatarURL != nil && *req.DefaultAvatarURL != "" {
		avatarSource := strings.TrimSpace(*req.DefaultAvatarURL)
		if !strings.Contains(avatarSource, "ik.imagekit.io") {
			fileName := fmt.Sprintf("avatar_%s.jpg", userID.String())
			ikURL, err := u.imagekit.UploadFromURL(ctx, avatarSource, fileName, "/avatars")
			if err == nil && ikURL != "" {
				avatarSource = ikURL
			}
		}
		user.DefaultAvatarURL = &avatarSource
	}

	if req.ActiveRole != nil && *req.ActiveRole != "" {
		user.ActiveRole = *req.ActiveRole
	}

	if req.PhoneNumber != nil && *req.PhoneNumber != "" {
		phone := *req.PhoneNumber
		found := false

		for i := range user.PhoneNumbers {
			if user.PhoneNumbers[i].Number == phone {
				user.PhoneNumbers[i].IsPrimary = true
				found = true
			} else {
				user.PhoneNumbers[i].IsPrimary = false
			}
		}

		if !found {
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
func (u *UserUsecase) GetPhoneNumbers(ctx context.Context, userID uuid.UUID, req dto.GetPhoneNumbersRequest) ([]dto.PhoneNumberItem, api.PaginationMeta, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return nil, api.PaginationMeta{}, errors.New(string(i18n.KeyUserNotFound))
	}

	req.SetDefaults()

	var filtered []entity.PhoneNumberItem
	for _, p := range user.PhoneNumbers {
		if req.Number != "" && !strings.Contains(strings.ToLower(p.Number), strings.ToLower(req.Number)) {
			continue
		}
		filtered = append(filtered, p)
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		if filtered[i].IsPrimary {
			return true
		}
		if filtered[j].IsPrimary {
			return false
		}
		return false
	})

	totalItems := len(filtered)
	totalPages := int(math.Ceil(float64(totalItems) / float64(req.Limit)))

	startIndex := (req.Page - 1) * req.Limit
	endIndex := startIndex + req.Limit

	if startIndex >= totalItems {
		meta := api.PaginationMeta{
			TotalItems:  totalItems,
			TotalPages:  totalPages,
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: false,
			HasPrevPage: req.Page > 1,
		}
		return []dto.PhoneNumberItem{}, meta, nil
	}

	if endIndex > totalItems {
		endIndex = totalItems
	}

	pagedItems := filtered[startIndex:endIndex]
	items := make([]dto.PhoneNumberItem, 0, len(pagedItems))
	for _, p := range pagedItems {
		items = append(items, dto.PhoneNumberItem{
			Number:    p.Number,
			IsPrimary: p.IsPrimary,
			Roles:     p.Roles,
		})
	}

	meta := api.PaginationMeta{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: req.Page < totalPages,
		HasPrevPage: req.Page > 1,
	}

	return items, meta, nil
}

// Replace/Merge AddPhoneNumber, UpdatePhoneNumber, & DeletePhoneNumber menjadi SyncPhoneNumbers
func (u *UserUsecase) SyncPhoneNumbers(ctx context.Context, userID uuid.UUID, req dto.SavePhoneNumbersRequest) error {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}

	// 1. Cek Apakah Kosong
	if len(req.PhoneNumbers) == 0 {
		return errors.New(string(i18n.KeyUserPhoneCannotBeEmpty))
	}

	primaryCount := 0
	seenNumbers := make(map[string]bool)
	newPhoneList := make(entity.PhoneNumbers, 0, len(req.PhoneNumbers))

	for _, item := range req.PhoneNumbers {
		cleanNum := strings.TrimSpace(item.Number)
		if cleanNum == "" {
			continue
		}

		// Cek Duplikasi Nomor di dalam Payload JSON
		if seenNumbers[cleanNum] {
			return errors.New(string(i18n.KeyUserPhoneDuplicate))
		}
		seenNumbers[cleanNum] = true

		if item.IsPrimary {
			primaryCount++
		}

		newPhoneList = append(newPhoneList, entity.PhoneNumberItem{
			Number:    cleanNum,
			IsPrimary: item.IsPrimary,
			Roles:     item.Roles,
		})
	}

	// 2. Cek Apakah Tidak Ada Primary Sama Sekali
	if primaryCount == 0 {
		return errors.New(string(i18n.KeyUserPhonePrimaryRequired))
	}

	// 3. Cek Apakah Primary Lebih Dari 1
	if primaryCount > 1 {
		return errors.New(string(i18n.KeyUserPhoneMultiplePrimaryNotAllowed))
	}

	user.PhoneNumbers = newPhoneList
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

	avatarURL := req.AvatarURL
	if avatarURL != "" && !strings.Contains(avatarURL, "ik.imagekit.io") {
		fileName := fmt.Sprintf("persona_%s_%s.jpg", role, userID.String())
		ikURL, err := u.imagekit.UploadFromURL(ctx, avatarURL, fileName, "/personas")
		if err == nil && ikURL != "" {
			avatarURL = ikURL
		}
	}

	user.RoleProfiles[role] = entity.RoleProfileItem{
		DisplayName: req.DisplayName,
		AvatarURL:   avatarURL,
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return dto.PersonaProfileResponse{}, err
	}

	return u.GetPersonaProfile(ctx, userID, role)
}

// 5. Document Upload & Watermarking
func (u *UserUsecase) GetDocument(ctx context.Context, userID uuid.UUID, req dto.GetDocumentRequest) ([]dto.GetDocumentResponse, api.PaginationMeta, error) {
	req.SetDefaults()

	docs, totalItems, err := u.repo.GetDocument(ctx, userID, req.Name, req.NIK, req.Page, req.Limit)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(req.Limit)))

	res := make([]dto.GetDocumentResponse, 0, len(docs))
	for _, d := range docs {
		domicile := ""
		if d.DomicileCity != nil {
			domicile = *d.DomicileCity
		}
		res = append(res, dto.GetDocumentResponse{
			ID:           d.ID.String(),
			FullNameKTP:  d.FullNameKTP,
			NIK:          d.NIK,
			KTPPhotoURL:  d.KTPPhotoURL,
			DomicileCity: domicile,
		})
	}

	meta := api.PaginationMeta{
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: req.Page < totalPages,
		HasPrevPage: req.Page > 1,
	}

	return res, meta, nil
}

func (u *UserUsecase) UploadDocument(ctx context.Context, userID uuid.UUID, req dto.UploadDocumentRequest) error {
	existingIdentity, err := u.repo.FindIdentityByNIK(ctx, req.NIK)
	if err != nil {
		return err
	}
	if existingIdentity != nil && existingIdentity.UserID != userID {
		return errors.New(string(i18n.KeyUserDocumentNIKExists))
	}

	rawBase64 := req.KTPPhoto
	if idx := strings.Index(rawBase64, ","); idx != -1 {
		rawBase64 = rawBase64[idx+1:]
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(rawBase64)
	if err != nil {
		return errors.New(string(i18n.KeyUserDocumentFileInvalid))
	}

	purpose := storage.PurposeStallVerification
	watermarkedBytes, err := storage.ApplyWatermarkFromBytes(decodedBytes, purpose)
	if err != nil {
		return errors.New(string(i18n.KeyUserDocumentWatermarkFailed))
	}

	watermarkedBase64 := base64.StdEncoding.EncodeToString(watermarkedBytes)
	fileName := fmt.Sprintf("ktp_%s.png", userID.String())

	uploadedURL, err := u.imagekit.UploadFromURL(ctx, watermarkedBase64, fileName, "/users/identity_documents")
	if err != nil {
		return errors.New(string(i18n.KeyUserDocumentFailedToUpload))
	}

	domicile := req.DomicileCity
	identity := &entity.UserIdentityProfile{
		UserID:       userID,
		FullNameKTP:  req.FullNameKTP,
		NIK:          req.NIK,
		KTPPhotoURL:  uploadedURL,
		DomicileCity: &domicile,
	}

	return u.repo.UpsertIdentityProfile(ctx, identity)
}

func (u *UserUsecase) DeleteDocument(ctx context.Context, userID uuid.UUID) error {
	return u.repo.DeleteIdentityProfile(ctx, userID)
}
