package usecase

import (
	"context"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/wallet/dto"
	"lapakita-backend/pkg/api"

	"github.com/google/uuid"
)

func (u *WalletUsecase) GetPayoutMethods(ctx context.Context, userID uuid.UUID, req dto.GetPayoutMethodsRequest) ([]dto.PayoutMethodResponse, api.PaginationMeta, error) {
	methods, meta, err := u.Repo.GetByUserID(ctx, userID, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	responses := make([]dto.PayoutMethodResponse, 0, len(methods))
	for _, method := range methods {
		responses = append(responses, toPayoutMethodResponse(method))
	}
	return responses, meta, nil
}

func toPayoutMethodResponse(method entity.PayoutMethod) dto.PayoutMethodResponse {
	return dto.PayoutMethodResponse{
		ID:                api.EncodeResponse(method.ID),
		PayoutType:        method.PayoutType,
		ProviderCode:      method.ProviderCode,
		ProviderName:      method.ProviderName,
		AccountNumber:     method.AccountNumber,
		AccountHolderName: method.AccountHolderName,
		IsPrimary:         method.IsPrimary,
		CreatedAt:         method.CreatedAt,
		UpdatedAt:         method.UpdatedAt,
	}
}
