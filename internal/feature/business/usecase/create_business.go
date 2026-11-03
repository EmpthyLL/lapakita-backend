package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) Create(ctx context.Context, req dto.CreateBusinessRequest) (*dto.CreateBusinessResponse, error) {
	businessTypeUUID, err := api.DecodeRequest(req.BusinessTypeID)
	if err != nil {
		return nil, api.ErrInvalidID
	}

	var logoURL *string
	if req.LogoURL != nil && strings.TrimSpace(*req.LogoURL) != "" {
		sourceLogo := strings.TrimSpace(*req.LogoURL)
		fileName := fmt.Sprintf("business_logo_%s_%d.jpg", req.UserID.String(), time.Now().Unix())

		// Langsung di-upload ke SeaweedFS tanpa watermark
		seaweedURL, err := u.seaweed.UploadFromURL(ctx, sourceLogo, fileName, "/logos")
		if err == nil && seaweedURL != "" {
			logoURL = &seaweedURL
		} else {
			logoURL = &sourceLogo
		}
	}

	business := &entity.Business{
		UserID:         req.UserID,
		BusinessTypeID: businessTypeUUID,
		Name:           req.Name,
		LogoURL:        logoURL,
	}

	if err := u.businessRepo.Create(ctx, business); err != nil {
		return nil, err
	}

	return &dto.CreateBusinessResponse{
		ID:             api.EncodeResponse(business.ID),
		UserID:         api.EncodeResponse(business.UserID),
		BusinessTypeID: api.EncodeResponse(business.BusinessTypeID),
		Name:           business.Name,
		LogoURL:        business.LogoURL,
		CreatedAt:      business.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
