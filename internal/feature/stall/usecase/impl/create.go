package impl

import (
	"context"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/stall/dto"
	usecase "lapakita-backend/internal/feature/stall/usecase"

	"github.com/google/uuid"
)

func Create(ctx context.Context, u *usecase.StallUsecase, ownerID uuid.UUID, req dto.CreateStallRequest) (dto.StallDetailResponse, error) {
	displayMedia, err := usecase.ProcessDisplayMedia(ctx, u.SeaweedFS, req.DisplayMedia)
	if err != nil {
		return dto.StallDetailResponse{}, err
	}
	legalDocs, err := usecase.ProcessLegalDocs(ctx, u.SeaweedFS, req.LegalDocuments)
	if err != nil {
		return dto.StallDetailResponse{}, err
	}

	var opHours *entity.OperatingHours
	if req.OperatingHours != nil {
		opHours = &entity.OperatingHours{OpeningTime: req.OperatingHours.OpeningTime, ClosingTime: req.OperatingHours.ClosingTime, Is24Hours: req.OperatingHours.Is24Hours}
	}
	var eventSched *entity.EventSchedule
	if req.EventSchedule != nil {
		eventSched = &entity.EventSchedule{EventName: req.EventSchedule.EventName, StartDate: req.EventSchedule.StartDate, EndDate: req.EventSchedule.EndDate, RegistrationDeadline: req.EventSchedule.RegistrationDeadline}
	}
	var slotInf *entity.SlotInfo
	if req.SlotInfo != nil {
		slotInf = &entity.SlotInfo{TotalSlots: req.SlotInfo.TotalSlots, AvailableSlots: req.SlotInfo.AvailableSlots}
	}

	country := req.Country
	if country == "" {
		country = "Indonesia"
	}
	countryCode := req.CountryCode
	if countryCode == "" {
		countryCode = "ID"
	}

	stall := &entity.Stall{
		StallOwnerID: ownerID, Title: req.Title, Description: req.Description, PropertyType: req.PropertyType,
		PermanenceType: entity.StallPermanenceType(req.PermanenceType), Placement: entity.StallPlacementType(req.Placement),
		SizeSqm: req.SizeSqm, LengthMeters: req.LengthMeters, WidthMeters: req.WidthMeters, FloorLevel: req.FloorLevel,
		ElectricityCapacityVA: req.ElectricityCapacityVA, ParentComplexName: req.ParentComplexName, OperatingHours: opHours,
		EventSchedule: eventSched, SlotInfo: slotInf, StreetAddress: req.StreetAddress, Suburb: req.Suburb,
		District: req.District, City: req.City, Province: req.Province, Country: country, CountryCode: countryCode,
		PostalCode: req.PostalCode, Latitude: req.Latitude, Longitude: req.Longitude, MapURL: req.MapURL,
		EmbeddedMapURL: req.EmbeddedMapURL, AllowedPaymentCycles: req.AllowedPaymentCycles, DailyRate: req.DailyRate,
		MonthlyRate: req.MonthlyRate, QuarterlyRate: req.QuarterlyRate, SemesterlyRate: req.SemesterlyRate,
		YearlyRate: req.YearlyRate, SecurityDeposit: req.SecurityDeposit, MinimumLeaseMonths: req.MinimumLeaseMonths,
		MinimumLeaseDays: req.MinimumLeaseDays, StartDateOptions: req.StartDateOptions, UtilityTerms: req.UtilityTerms,
		FacilityValues: req.FacilityValues, AllowedBusinessTypeIDs: req.AllowedBusinessTypeIDs, HouseRules: req.HouseRules,
		DisplayMedia: displayMedia, LegalDocuments: legalDocs, IsPublished: req.IsPublished,
	}
	if len(req.NearbyLandmarks) > 0 {
		landmarks := make(entity.NearbyLandmarkArray, 0, len(req.NearbyLandmarks))
		for _, lm := range req.NearbyLandmarks {
			landmarks = append(landmarks, entity.NearbyLandmarkItem{Name: lm, DistanceKm: 0.5, CategoryValue: "general"})
		}
		stall.NearbyLandmarks = landmarks
	}
	if req.EventOperatingDays != nil {
		val := entity.EventOperatingDaysType(*req.EventOperatingDays)
		stall.EventOperatingDays = &val
	}
	if req.EventAttendanceReq != nil {
		val := entity.AttendanceRequirementType(*req.EventAttendanceReq)
		stall.EventAttendanceReq = &val
	}
	if req.EventCancellationPolicy != nil {
		val := entity.CancellationPolicyType(*req.EventCancellationPolicy)
		stall.EventCancellationPolicy = &val
	}
	if err := u.StallRepository.Create(ctx, stall); err != nil {
		return dto.StallDetailResponse{}, err
	}
	return GetByID(ctx, u, stall.ID)
}
