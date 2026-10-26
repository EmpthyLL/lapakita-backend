package usecase

import (
	"context"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/pkg/api"
)

func (u *MasterUsecase) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	return u.MasterClient.SearchGeneral(ctx, req)
}
