package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) Update(ctx context.Context, req dto.UpdateBusinessRequest) (*dto.UpdateBusinessResponse, error) {
	businessUUID, err := api.DecodeRequest(req.ID)
	if err != nil {
		return nil, api.ErrInvalidID
	}

	business, err := u.businessRepo.FindByID(ctx, businessUUID)
	if err != nil {
		return nil, err
	}
	if business == nil {
		return nil, errors.New("business not found")
	}

	if req.BusinessTypeID != nil {
		btUUID, err := api.DecodeRequest(*req.BusinessTypeID)
		if err != nil {
			return nil, api.ErrInvalidID
		}
		business.BusinessTypeID = btUUID
	}

	if req.Name != nil {
		business.Name = *req.Name
	}

	if req.LogoURL != nil {
		if strings.TrimSpace(*req.LogoURL) == "" {
			business.LogoURL = nil
		} else {
			sourceLogo := strings.TrimSpace(*req.LogoURL)
			fileName := fmt.Sprintf("business_logo_%s_%d.jpg", business.ID.String(), time.Now().Unix())

			// Langsung di-upload ke SeaweedFS tanpa watermark
			seaweedURL, err := u.seaweed.UploadFromURL(ctx, sourceLogo, fileName, "/logos")
			if err == nil && seaweedURL != "" {
				business.LogoURL = &seaweedURL
			} else {
				business.LogoURL = &sourceLogo
			}
		}
	}

	if err := u.businessRepo.Update(ctx, business); err != nil {
		return nil, err
	}

	return &dto.UpdateBusinessResponse{
		ID:             api.EncodeResponse(business.ID),
		BusinessTypeID: api.EncodeResponse(business.BusinessTypeID),
		Name:           business.Name,
		LogoURL:        business.LogoURL,
		UpdatedAt:      business.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
