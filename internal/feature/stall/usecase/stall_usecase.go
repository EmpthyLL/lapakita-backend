package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/stall/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/storage"
	"lapakita-backend/pkg/utils"

	"github.com/google/uuid"
)

func UploadMediaToSeaweedFS(ctx context.Context, seaweed *storage.SeaweedFSService, source string, folder string) (string, error) {
	cleanSource := strings.TrimSpace(source)
	if cleanSource == "" {
		return "", nil
	}
	if seaweed.IsStoredURL(cleanSource) {
		return cleanSource, nil
	}

	fileName := fmt.Sprintf("stall_%s.jpg", uuid.New().String())
	seaweedURL, err := seaweed.UploadFromURL(ctx, cleanSource, fileName, folder)
	if err != nil {
		return "", err
	}
	return seaweedURL, nil
}

func ProcessDisplayMedia(ctx context.Context, seaweed *storage.SeaweedFSService, media dto.DisplayMediaDTO) (entity.DisplayMedia, error) {
	mainImg, err := UploadMediaToSeaweedFS(ctx, seaweed, media.MainImage, "/stalls/main")
	if err != nil {
		return entity.DisplayMedia{}, err
	}

	facilityImgs := make([]entity.FacilityImage, 0, len(media.FacilityImages))
	for _, img := range media.FacilityImages {
		uploaded, err := UploadMediaToSeaweedFS(ctx, seaweed, img, "/stalls/facilities")
		if err != nil {
			return entity.DisplayMedia{}, err
		}
		if uploaded != "" {
			facilityImgs = append(facilityImgs, entity.FacilityImage{URL: uploaded})
		}
	}

	return entity.DisplayMedia{
		MainImage:      mainImg,
		FacilityImages: facilityImgs,
	}, nil
}

func ProcessLegalDocs(ctx context.Context, seaweed *storage.SeaweedFSService, docs []string) ([]string, error) {
	uploadedDocs := make([]string, 0, len(docs))
	for _, doc := range docs {
		uploaded, err := UploadMediaToSeaweedFS(ctx, seaweed, doc, "/stalls/documents")
		if err != nil {
			return nil, err
		}
		if uploaded != "" {
			uploadedDocs = append(uploadedDocs, uploaded)
		}
	}
	return uploadedDocs, nil
}

func DaysUntilDate(dateValue string) int {
	deadline, err := time.ParseInLocation("2006-01-02", dateValue, time.Local)
	if err != nil {
		return 0
	}
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return int(deadline.Sub(today).Hours() / 24)
}

func MapToCompactResponse(s *entity.Stall) dto.CompactStallResponse {
	areaStr := s.StreetAddress
	if s.District != nil && *s.District != "" {
		areaStr = *s.District
	} else if s.Suburb != nil && *s.Suburb != "" {
		areaStr = *s.Suburb
	}

	locationSummary := dto.StallLocationSummary{
		Area:        areaStr,
		City:        s.City,
		CountryCode: s.CountryCode,
	}

	cheapestVal, periodKey := CalculateCheapestPrice(s)
	priceFormatted := fmt.Sprintf("Rp %s", utils.FormatRupiah(cheapestVal))

	base := dto.BaseStallResponse{
		ID:                     api.EncodeResponse(s.ID),
		Title:                  s.Title,
		ImageURL:               s.DisplayMedia.MainImage,
		Location:               locationSummary,
		PropertyType:           s.PropertyType,
		Placement:              string(s.Placement),
		CheapestPriceFormatted: priceFormatted,
		CheapestPricePeriod:    periodKey,
		Rating:                 s.RatingAvg,
		ReviewCount:            s.ReviewCount,
	}

	switch s.PermanenceType {
	case entity.StallPermanenceSemi:
		openTime, closeTime := "10:00", "22:00"
		if s.OperatingHours != nil {
			openTime = s.OperatingHours.OpeningTime
			closeTime = s.OperatingHours.ClosingTime
		}

		res := dto.SemiPermanentStallResponse{
			BaseStallResponse: base,
			PermanenceType:    "semi-permanent",
		}
		res.OperatingHours.Open = openTime
		res.OperatingHours.Close = closeTime
		return dto.CompactStallResponse{SemiPermanentStallResponse: &res}

	case entity.StallPermanenceTemporary:
		registrationDeadlineDays := 0
		if s.EventSchedule != nil {
			registrationDeadlineDays = DaysUntilDate(s.EventSchedule.RegistrationDeadline)
		}

		durationDays := 1
		if s.EventSchedule != nil && s.EventSchedule.StartDate != "" && s.EventSchedule.EndDate != "" {
			t1, err1 := time.Parse("2006-01-02", s.EventSchedule.StartDate)
			t2, err2 := time.Parse("2006-01-02", s.EventSchedule.EndDate)
			if err1 == nil && err2 == nil {
				durationDays = int(t2.Sub(t1).Hours()/24) + 1
			}
		}

		res := dto.TemporaryStallResponse{
			BaseStallResponse: base,
			PermanenceType:    "temporary",
		}
		res.Event.RegistrationDeadlineDays = registrationDeadlineDays
		res.Event.DurationDays = durationDays
		return dto.CompactStallResponse{TemporaryStallResponse: &res}

	default:
		sqm := 0.0
		if s.SizeSqm != nil {
			sqm = *s.SizeSqm
		}
		floors := 1
		if s.FloorLevel != nil {
			floors = *s.FloorLevel
		}

		res := dto.PermanentStallResponse{
			BaseStallResponse: base,
			PermanenceType:    "permanent",
		}
		res.Space.SizeSqm = sqm
		res.Space.FloorCount = floors
		return dto.CompactStallResponse{PermanentStallResponse: &res}
	}
}

func CalculateCheapestPrice(s *entity.Stall) (float64, string) {
	minPrice := 0.0
	period := "month"

	if s.DailyRate != nil && *s.DailyRate > 0 {
		minPrice = *s.DailyRate
		period = "day"
	}
	if s.MonthlyRate != nil && *s.MonthlyRate > 0 && (minPrice == 0 || *s.MonthlyRate < minPrice) {
		minPrice = *s.MonthlyRate
		period = "month"
	}
	if s.QuarterlyRate != nil && *s.QuarterlyRate > 0 && (minPrice == 0 || *s.QuarterlyRate < minPrice) {
		minPrice = *s.QuarterlyRate
		period = "quarter"
	}
	if s.SemesterlyRate != nil && *s.SemesterlyRate > 0 && (minPrice == 0 || *s.SemesterlyRate < minPrice) {
		minPrice = *s.SemesterlyRate
		period = "semester"
	}
	if s.YearlyRate != nil && *s.YearlyRate > 0 && (minPrice == 0 || *s.YearlyRate < minPrice) {
		minPrice = *s.YearlyRate
		period = "year"
	}

	return minPrice, period
}

func MapToDetailResponse(s *entity.Stall, owner *entity.User) dto.StallDetailResponse {
	ownerSummary := dto.OwnerProfileSummary{
		ID:          api.EncodeResponse(s.StallOwnerID),
		Name:        "Unknown Owner",
		Contact:     "",
		AvatarURL:   "",
		Rating:      5.0,
		ReviewCount: 0,
		JoinedYear:  fmt.Sprintf("%d", time.Now().Year()),
	}

	if owner != nil {
		ownerSummary.Name = owner.Name
		ownerSummary.ContactDialCode = owner.PhoneNumbers.GetDialCodeForRole("owner")
		ownerSummary.Contact = owner.PhoneNumbers.GetNumberForRole("owner")

		avatarURL := ""
		if profile, exists := owner.RoleProfiles["owner"]; exists && profile.AvatarURL != "" {
			avatarURL = profile.AvatarURL
		} else if owner.DefaultAvatarURL != nil {
			avatarURL = *owner.DefaultAvatarURL
		}
		ownerSummary.AvatarURL = avatarURL
		ownerSummary.JoinedYear = fmt.Sprintf("%d", owner.CreatedAt.Year())
	}

	facilityImgs := make([]dto.FacilityImageDTO, 0, len(s.DisplayMedia.FacilityImages))
	for idx, imgVal := range s.DisplayMedia.FacilityImages {
		imgStr := imgVal.URL

		facilityImgs = append(facilityImgs, dto.FacilityImageDTO{
			ID:      fmt.Sprintf("img-%d", idx+1),
			URL:     imgStr,
			Caption: fmt.Sprintf("Facility Image %d", idx+1),
		})
	}

	virtualTour := ""
	if s.DisplayMedia.VirtualTour360URL != nil {
		virtualTour = *s.DisplayMedia.VirtualTour360URL
	}

	mediaDTO := dto.StallMediaDTO{
		MainImage:         s.DisplayMedia.MainImage,
		FacilityImages:    facilityImgs,
		VirtualTour360URL: &virtualTour,
	}

	suburb := ""
	if s.Suburb != nil {
		suburb = *s.Suburb
	}
	district := ""
	if s.District != nil {
		district = *s.District
	}
	postalCode := ""
	if s.PostalCode != nil {
		postalCode = *s.PostalCode
	}
	mapURL := ""
	if s.MapURL != nil {
		mapURL = *s.MapURL
	}
	embeddedMapURL := ""
	if s.EmbeddedMapURL != nil {
		embeddedMapURL = *s.EmbeddedMapURL
	}

	addressDTO := dto.StallAddressDTO{
		Street:         s.StreetAddress,
		Suburb:         suburb,
		District:       district,
		City:           s.City,
		Country:        s.Country,
		CountryCode:    s.CountryCode,
		Province:       s.Province,
		PostalCode:     postalCode,
		MapURL:         mapURL,
		EmbeddedMapURL: embeddedMapURL,
	}

	nearbyLandmarks := make([]dto.NearbyLandmarkDTO, 0, len(s.NearbyLandmarks))
	for _, l := range s.NearbyLandmarks {
		nearbyLandmarks = append(nearbyLandmarks, dto.NearbyLandmarkDTO{
			CategoryValue: l.CategoryValue,
			Name:          l.Name,
			DistanceKm:    l.DistanceKm,
		})
	}

	pricingDTO := dto.MultiPeriodPricing{
		DailyRate:            s.DailyRate,
		MonthlyRate:          s.MonthlyRate,
		QuarterlyRate:        s.QuarterlyRate,
		SemesterlyRate:       s.SemesterlyRate,
		YearlyRate:           s.YearlyRate,
		SecurityDeposit:      s.SecurityDeposit,
		AllowedPaymentCycles: s.AllowedPaymentCycles,
	}

	desc := ""
	if s.Description != nil {
		desc = *s.Description
	}

	utilityTerms := ""
	if s.UtilityTerms != nil {
		utilityTerms = *s.UtilityTerms
	}

	switch s.PermanenceType {
	case entity.StallPermanenceSemi:
		parentComplex := ""
		if s.ParentComplexName != nil {
			parentComplex = *s.ParentComplexName
		}
		opHours := dto.StallOperatingHoursDetail{}
		if s.OperatingHours != nil {
			opHours.OpeningTime = s.OperatingHours.OpeningTime
			opHours.ClosingTime = s.OperatingHours.ClosingTime
			opHours.Is24Hours = s.OperatingHours.Is24Hours
		}

		minMonths := 1
		if s.MinimumLeaseMonths != nil {
			minMonths = *s.MinimumLeaseMonths
		}

		res := dto.SemiPermanentStallDetail{
			ID:                    api.EncodeResponse(s.ID),
			Title:                 s.Title,
			Description:           desc,
			Media:                 mediaDTO,
			PropertyType:          s.PropertyType,
			PropertyTypeValue:     s.PropertyType,
			Placement:             string(s.Placement),
			ElectricityCapacityVA: utils.GetIntVal(s.ElectricityCapacityVA),
			Address:               addressDTO,
			NearbyLandmarks:       nearbyLandmarks,
			Pricing:               pricingDTO,
			FacilityValues:        s.FacilityValues,
			HouseRules:            s.HouseRules,
			Rating:                s.RatingAvg,
			ReviewCount:           s.ReviewCount,
			Owner:                 ownerSummary,
			PermanenceType:        "semi-permanent",
			ParentComplexName:     parentComplex,
			OperatingHours:        opHours,
			LeaseRules: dto.PermanentLeaseRules{
				MinimumLeaseMonths: minMonths,
				StartDateOptions:   s.StartDateOptions,
				UtilityTerms:       utilityTerms,
			},
		}
		return dto.StallDetailResponse{SemiPermanentStallDetail: &res}

	case entity.StallPermanenceTemporary:
		eventMeta := dto.StallEventMeta{}

		if s.EventSchedule != nil {
			eventMeta.EventName = s.EventSchedule.EventName
			eventMeta.EventStartDate = s.EventSchedule.StartDate
			eventMeta.EventEndDate = s.EventSchedule.EndDate
			eventMeta.RegistrationDeadlineDaysBefore = DaysUntilDate(s.EventSchedule.RegistrationDeadline)
		}
		if s.SlotInfo != nil {
			eventMeta.TotalSlots = s.SlotInfo.TotalSlots
			eventMeta.AvailableSlots = s.SlotInfo.AvailableSlots
		}

		minDays := 1
		if s.MinimumLeaseDays != nil {
			minDays = *s.MinimumLeaseDays
		}

		opDays := ""
		if s.EventOperatingDays != nil {
			opDays = string(*s.EventOperatingDays)
		}
		attReq := ""
		if s.EventAttendanceReq != nil {
			attReq = string(*s.EventAttendanceReq)
		}
		cancelPol := ""
		if s.EventCancellationPolicy != nil {
			cancelPol = string(*s.EventCancellationPolicy)
		}

		res := dto.TemporaryStallDetail{
			ID:                    api.EncodeResponse(s.ID),
			Title:                 s.Title,
			Description:           desc,
			Media:                 mediaDTO,
			PropertyType:          s.PropertyType,
			PropertyTypeValue:     s.PropertyType,
			Placement:             string(s.Placement),
			ElectricityCapacityVA: utils.GetIntVal(s.ElectricityCapacityVA),
			Address:               addressDTO,
			NearbyLandmarks:       nearbyLandmarks,
			Pricing:               pricingDTO,
			FacilityValues:        s.FacilityValues,
			HouseRules:            s.HouseRules,
			Rating:                s.RatingAvg,
			ReviewCount:           s.ReviewCount,
			Owner:                 ownerSummary,
			PermanenceType:        "temporary",
			EventMeta:             eventMeta,
			LeaseRules: dto.TemporaryLeaseRules{
				MinimumLeaseDays:      minDays,
				StartDateOptions:      s.StartDateOptions,
				UtilityTerms:          utilityTerms,
				OperatingDays:         opDays,
				AttendanceRequirement: attReq,
				CancellationPolicy:    cancelPol,
			},
		}
		return dto.StallDetailResponse{TemporaryStallDetail: &res}

	default:
		minMonths := 1
		if s.MinimumLeaseMonths != nil {
			minMonths = *s.MinimumLeaseMonths
		}

		res := dto.PermanentStallDetail{
			ID:                    api.EncodeResponse(s.ID),
			Title:                 s.Title,
			Description:           desc,
			Media:                 mediaDTO,
			PropertyType:          s.PropertyType,
			PropertyTypeValue:     s.PropertyType,
			Placement:             string(s.Placement),
			ElectricityCapacityVA: utils.GetIntVal(s.ElectricityCapacityVA),
			Address:               addressDTO,
			NearbyLandmarks:       nearbyLandmarks,
			Pricing:               pricingDTO,
			FacilityValues:        s.FacilityValues,
			HouseRules:            s.HouseRules,
			Rating:                s.RatingAvg,
			ReviewCount:           s.ReviewCount,
			Owner:                 ownerSummary,
			PermanenceType:        "permanent",
			SizeSqm:               utils.GetFloatVal(s.SizeSqm),
			Dimensions: struct {
				LengthMeters float64 `json:"lengthMeters"`
				WidthMeters  float64 `json:"widthMeters"`
			}{
				LengthMeters: utils.GetFloatVal(s.LengthMeters),
				WidthMeters:  utils.GetFloatVal(s.WidthMeters),
			},
			FloorLevel: utils.GetIntVal(s.FloorLevel),
			LeaseRules: dto.PermanentLeaseRules{
				MinimumLeaseMonths: minMonths,
				StartDateOptions:   s.StartDateOptions,
				UtilityTerms:       utilityTerms,
			},
		}
		return dto.StallDetailResponse{PermanentStallDetail: &res}
	}
}
