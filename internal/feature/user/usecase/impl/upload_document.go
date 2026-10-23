package impl

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	usermessages "lapakita-backend/internal/feature/user/messages"
	"strings"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/i18n"
	"lapakita-backend/pkg/storage"

	"github.com/google/uuid"
)

func UploadDocument(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, req dto.UploadDocumentRequest) error {
	existingUserDoc, err := u.Repo.FindIdentityByUserIDAndDocNumber(ctx, userID, req.DocumentNumber)
	if err != nil {
		return err
	}
	if existingUserDoc != nil {
		return errors.New(string(usermessages.KeyUserDocumentNumberAlreadyAdded))
	}
	existingGlobalDoc, err := u.Repo.FindIdentityByDocumentNumber(ctx, req.DocumentNumber)
	if err != nil {
		return err
	}
	if existingGlobalDoc != nil && existingGlobalDoc.UserID != userID {
		return errors.New(string(usermessages.KeyUserDocumentNIKExists))
	}
	rawBase64 := req.DocumentPhoto
	if idx := strings.Index(rawBase64, ","); idx != -1 {
		rawBase64 = rawBase64[idx+1:]
	}
	decodedBytes, err := base64.StdEncoding.DecodeString(rawBase64)
	if err != nil {
		return errors.New(string(usermessages.KeyUserDocumentFileInvalid))
	}
	watermarkedBytes, err := storage.ApplyWatermarkFromBytes(decodedBytes, storage.PurposeStallVerification)
	if err != nil {
		return errors.New(string(usermessages.KeyUserDocumentWatermarkFailed))
	}
	watermarkedBase64 := base64.StdEncoding.EncodeToString(watermarkedBytes)
	fileName := fmt.Sprintf("%s_%s_%d.jpg", req.DocumentType, userID.String(), time.Now().Unix())
	uploadedURL, err := u.Seaweed.UploadFromURL(ctx, watermarkedBase64, fileName, "/legals")
	if err != nil {
		return errors.New(string(usermessages.KeyUserDocumentFailedToUpload))
	}
	countryCode := req.CountryCode
	if countryCode == "" {
		countryCode = "ID"
	}
	docLabel := req.DocumentLabel
	if docLabel == "" {
		docLabel = "KTP"
	}
	identity := &entity.UserIdentityProfile{UserID: userID, CountryCode: countryCode, DocumentType: req.DocumentType, DocumentLabel: docLabel, FullNameIdentity: req.FullNameIdentity, DocumentNumber: req.DocumentNumber, DocumentPhotoURL: uploadedURL}
	return u.Repo.CreateIdentityProfile(ctx, identity)
}
