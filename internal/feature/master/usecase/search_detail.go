package usecase

import (
	"context"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/pkg/api"
)

func (u *MasterUsecase) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	return u.MasterClient.SearchDetail(ctx, req)
}
