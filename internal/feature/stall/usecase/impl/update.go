package impl

import (
	"context"
	"errors"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/stall/dto"
	usecase "lapakita-backend/internal/feature/stall/usecase"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func Update(ctx context.Context, u *usecase.StallUsecase, ownerID uuid.UUID, req dto.UpdateStallRequest) (dto.StallDetailResponse, error) {
	stallID, err := api.DecodeRequest(req.ID)
	if err != nil {
		return dto.StallDetailResponse{}, errors.New(string(i18n.KeyInvalidPayload))
	}
	stall, err := u.StallRepository.FindByID(ctx, stallID)
	if err != nil || stall == nil {
		return dto.StallDetailResponse{}, errors.New("stall not found")
	}
	if stall.StallOwnerID != ownerID {
		return dto.StallDetailResponse{}, errors.New("unauthorized stall update")
	}
	stall.Title = req.Title
	stall.Description = req.Description
	stall.PropertyType = req.PropertyType
	stall.PermanenceType = entity.StallPermanenceType(req.PermanenceType)
	stall.Placement = entity.StallPlacementType(req.Placement)
	stall.SizeSqm = req.SizeSqm
	stall.LengthMeters = req.LengthMeters
	stall.WidthMeters = req.WidthMeters
	stall.FloorLevel = req.FloorLevel
	stall.ElectricityCapacityVA = req.ElectricityCapacityVA
	stall.ParentComplexName = req.ParentComplexName
	stall.StreetAddress = req.StreetAddress
	stall.Suburb = req.Suburb
	stall.District = req.District
	stall.City = req.City
	stall.Province = req.Province
	stall.PostalCode = req.PostalCode
	stall.Latitude = req.Latitude
	stall.Longitude = req.Longitude
	stall.MapURL = req.MapURL
	stall.EmbeddedMapURL = req.EmbeddedMapURL
	stall.AllowedPaymentCycles = req.AllowedPaymentCycles
	stall.DailyRate = req.DailyRate
	stall.MonthlyRate = req.MonthlyRate
	stall.QuarterlyRate = req.QuarterlyRate
	stall.SemesterlyRate = req.SemesterlyRate
	stall.YearlyRate = req.YearlyRate
	stall.SecurityDeposit = req.SecurityDeposit
	stall.MinimumLeaseMonths = req.MinimumLeaseMonths
	stall.MinimumLeaseDays = req.MinimumLeaseDays
	stall.StartDateOptions = req.StartDateOptions
	stall.UtilityTerms = req.UtilityTerms
	stall.FacilityValues = req.FacilityValues
	stall.AllowedBusinessTypeIDs = req.AllowedBusinessTypeIDs
	stall.HouseRules = req.HouseRules
	stall.IsPublished = req.IsPublished
	stall.UpdatedAt = time.Now()
	if len(req.NearbyLandmarks) > 0 {
		landmarks := make(entity.NearbyLandmarkArray, 0, len(req.NearbyLandmarks))
		for _, lm := range req.NearbyLandmarks {
			landmarks = append(landmarks, entity.NearbyLandmarkItem{Name: lm, DistanceKm: 0.5, CategoryValue: "general"})
		}
		stall.NearbyLandmarks = landmarks
	}
	if req.DisplayMedia.MainImage != "" || len(req.DisplayMedia.FacilityImages) > 0 {
		stall.DisplayMedia, err = usecase.ProcessDisplayMedia(ctx, u.SeaweedFS, req.DisplayMedia)
		if err != nil {
			return dto.StallDetailResponse{}, err
		}
	}
	if len(req.LegalDocuments) > 0 {
		stall.LegalDocuments, err = usecase.ProcessLegalDocs(ctx, u.SeaweedFS, req.LegalDocuments)
		if err != nil {
			return dto.StallDetailResponse{}, err
		}
	}
	if err := u.StallRepository.Update(ctx, stall); err != nil {
		return dto.StallDetailResponse{}, err
	}
	return GetByID(ctx, u, stall.ID)
}
