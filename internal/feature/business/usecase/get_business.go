package usecase

import (
	"context"

	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) GetByID(ctx context.Context, req dto.GetBusinessByIDRequest, lang string) (*dto.GetBusinessResponse, error) {
	businessUUID, err := api.DecodeRequest(req.ID)
	if err != nil {
		return nil, api.ErrInvalidID
	}

	business, err := u.businessRepo.FindByIDWithPreload(ctx, businessUUID)
	if err != nil {
		return nil, err
	}
	if business == nil {
		return nil, nil
	}

	var btName string
	if business.BusinessType != nil {
		btName = business.BusinessType.LabelLang.ID
		if lang == "en" {
			btName = business.BusinessType.LabelLang.EN
		}
	}

	return &dto.GetBusinessResponse{
		ID:               api.EncodeResponse(business.ID),
		UserID:           api.EncodeResponse(business.UserID),
		Name:             business.Name,
		LogoURL:          business.LogoURL,
		BusinessTypeName: btName,
		CreatedAt:        business.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:        business.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
