package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"lapakita-backend/internal/entity"
	authRepo "lapakita-backend/internal/feature/auth/repository"
	"lapakita-backend/internal/feature/stall/dto"
	"lapakita-backend/internal/feature/stall/repository"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"
	"lapakita-backend/pkg/storage"

	"github.com/google/uuid"
)

type StallUsecase struct {
	repo     *repository.StallRepository
	userRepo *authRepo.AuthRepository
	imagekit *storage.ImageKitService
}

func NewStallUsecase(repo *repository.StallRepository, userRepo *authRepo.AuthRepository, imagekit *storage.ImageKitService) *StallUsecase {
	return &StallUsecase{
		repo:     repo,
		userRepo: userRepo,
		imagekit: imagekit,
	}
}

func (u *StallUsecase) uploadMediaToImageKit(ctx context.Context, source string, folder string) string {
	cleanSource := strings.TrimSpace(source)
	if cleanSource == "" {
		return ""
	}
	if strings.Contains(cleanSource, "ik.imagekit.io") {
		return cleanSource
	}

	fileName := fmt.Sprintf("stall_%s.jpg", uuid.New().String())
	ikURL, err := u.imagekit.UploadFromURL(ctx, cleanSource, fileName, folder)
	if err != nil {
		return cleanSource
	}
	return ikURL
}

func (u *StallUsecase) processDisplayMedia(ctx context.Context, media dto.DisplayMediaDTO) entity.DisplayMedia {
	mainImg := u.uploadMediaToImageKit(ctx, media.MainImage, "/stalls/main")

	facilityImgs := make([]any, 0, len(media.FacilityImages))
	for _, img := range media.FacilityImages {
		if uploaded := u.uploadMediaToImageKit(ctx, img, "/stalls/facilities"); uploaded != "" {
			facilityImgs = append(facilityImgs, uploaded)
		}
	}

	return entity.DisplayMedia{
		MainImage:      mainImg,
		FacilityImages: facilityImgs,
	}
}

func (u *StallUsecase) processLegalDocs(ctx context.Context, docs []string) []string {
	uploadedDocs := make([]string, 0, len(docs))
	for _, doc := range docs {
		if uploaded := u.uploadMediaToImageKit(ctx, doc, "/stalls/documents"); uploaded != "" {
			uploadedDocs = append(uploadedDocs, uploaded)
		}
	}
	return uploadedDocs
}

func (u *StallUsecase) Create(ctx context.Context, ownerID uuid.UUID, req dto.CreateStallRequest) (any, error) {
	displayMedia := u.processDisplayMedia(ctx, req.DisplayMedia)
	legalDocs := u.processLegalDocs(ctx, req.LegalDocuments)

	var opHours *entity.OperatingHours
	if req.OperatingHours != nil {
		opHours = &entity.OperatingHours{
			OpeningTime: req.OperatingHours.OpeningTime,
			ClosingTime: req.OperatingHours.ClosingTime,
			Is24Hours:   req.OperatingHours.Is24Hours,
		}
	}

	var eventSched *entity.EventSchedule
	if req.EventSchedule != nil {
		eventSched = &entity.EventSchedule{
			EventName:                req.EventSchedule.EventName,
			StartDate:                req.EventSchedule.StartDate,
			EndDate:                  req.EventSchedule.EndDate,
			RegistrationDeadlineDays: req.EventSchedule.RegistrationDeadlineDays,
		}
	}

	var slotInf *entity.SlotInfo
	if req.SlotInfo != nil {
		slotInf = &entity.SlotInfo{
			TotalSlots:     req.SlotInfo.TotalSlots,
			AvailableSlots: req.SlotInfo.AvailableSlots,
		}
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
		StallOwnerID:           ownerID,
		Title:                  req.Title,
		Description:            req.Description,
		PropertyType:           req.PropertyType,
		PermanenceType:         entity.StallPermanenceType(req.PermanenceType),
		Placement:              entity.StallPlacementType(req.Placement),
		SizeSqm:                req.SizeSqm,
		LengthMeters:           req.LengthMeters,
		WidthMeters:            req.WidthMeters,
		FloorLevel:             req.FloorLevel,
		ElectricityCapacityVA:  req.ElectricityCapacityVA,
		ParentComplexName:      req.ParentComplexName,
		OperatingHours:         opHours,
		EventSchedule:          eventSched,
		SlotInfo:               slotInf,
		StreetAddress:          req.StreetAddress,
		Suburb:                 req.Suburb,
		District:               req.District,
		City:                   req.City,
		Province:               req.Province,
		Country:                country,
		CountryCode:            countryCode,
		PostalCode:             req.PostalCode,
		Latitude:               req.Latitude,
		Longitude:              req.Longitude,
		MapURL:                 req.MapURL,
		EmbeddedMapURL:         req.EmbeddedMapURL,
		AllowedPaymentCycles:   req.AllowedPaymentCycles,
		DailyRate:              req.DailyRate,
		MonthlyRate:            req.MonthlyRate,
		QuarterlyRate:          req.QuarterlyRate,
		SemesterlyRate:         req.SemesterlyRate,
		YearlyRate:             req.YearlyRate,
		SecurityDeposit:        req.SecurityDeposit,
		MinimumLeaseMonths:     req.MinimumLeaseMonths,
		MinimumLeaseDays:       req.MinimumLeaseDays,
		StartDateOptions:       req.StartDateOptions,
		UtilityTerms:           req.UtilityTerms,
		FacilityValues:         req.FacilityValues,
		AllowedBusinessTypeIDs: req.AllowedBusinessTypeIDs,
		HouseRules:             req.HouseRules,
		DisplayMedia:           displayMedia,
		LegalDocuments:         legalDocs,
		IsPublished:            req.IsPublished,
	}

	// Dynamic Nearby Landmark Items Mapping
	if len(req.NearbyLandmarks) > 0 {
		landmarks := make(entity.NearbyLandmarkArray, 0, len(req.NearbyLandmarks))
		for _, lm := range req.NearbyLandmarks {
			landmarks = append(landmarks, entity.NearbyLandmarkItem{
				Name:          lm,
				DistanceKm:    0.5,
				CategoryValue: "general",
			})
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

	if err := u.repo.Create(ctx, stall); err != nil {
		return nil, err
	}

	return u.GetByID(ctx, stall.ID)
}

func (u *StallUsecase) Update(ctx context.Context, ownerID uuid.UUID, req dto.UpdateStallRequest) (any, error) {
	stallID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, errors.New(string(i18n.KeyInvalidPayload))
	}

	stall, err := u.repo.FindByID(ctx, stallID)
	if err != nil || stall == nil {
		return nil, errors.New("stall not found")
	}

	if stall.StallOwnerID != ownerID {
		return nil, errors.New("unauthorized stall update")
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
			landmarks = append(landmarks, entity.NearbyLandmarkItem{
				Name:          lm,
				DistanceKm:    0.5,
				CategoryValue: "general",
			})
		}
		stall.NearbyLandmarks = landmarks
	}

	if req.DisplayMedia.MainImage != "" || len(req.DisplayMedia.FacilityImages) > 0 {
		stall.DisplayMedia = u.processDisplayMedia(ctx, req.DisplayMedia)
	}

	if len(req.LegalDocuments) > 0 {
		stall.LegalDocuments = u.processLegalDocs(ctx, req.LegalDocuments)
	}

	if err := u.repo.Update(ctx, stall); err != nil {
		return nil, err
	}

	return u.GetByID(ctx, stall.ID)
}

func (u *StallUsecase) GetByID(ctx context.Context, id uuid.UUID) (any, error) {
	stall, err := u.repo.FindByID(ctx, id)
	if err != nil || stall == nil {
		return nil, errors.New("stall not found")
	}

	ownerUser, _ := u.userRepo.FindUserByID(ctx, stall.StallOwnerID)
	return u.mapToDetailResponse(stall, ownerUser), nil
}

func (u *StallUsecase) Delete(ctx context.Context, ownerID uuid.UUID, id uuid.UUID) error {
	stall, err := u.repo.FindByID(ctx, id)
	if err != nil || stall == nil {
		return errors.New("stall not found")
	}
	if stall.StallOwnerID != ownerID {
		return errors.New("unauthorized delete action")
	}
	return u.repo.Delete(ctx, id)
}

func (u *StallUsecase) Search(ctx context.Context, req dto.SearchStallRequest) ([]dto.StallResponseData, api.PaginationMeta, error) {
	stalls, total, err := u.repo.Search(ctx, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	meta := api.PaginationMeta{
		TotalItems:  int(total),
		TotalPages:  totalPages,
		CurrentPage: page,
		PerPage:     limit,
		HasNextPage: page < totalPages,
		HasPrevPage: page > 1,
	}

	responseList := make([]dto.StallResponseData, 0, len(stalls))
	for _, s := range stalls {
		responseList = append(responseList, *u.mapToSimpleResponse(&s))
	}

	return responseList, meta, nil
}

func (u *StallUsecase) GetByOwnerID(ctx context.Context, req dto.GetOwnerStallsRequest) ([]dto.StallResponseData, api.PaginationMeta, error) {
	stalls, total, err := u.repo.FindByOwnerID(ctx, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	meta := api.PaginationMeta{
		TotalItems:  int(total),
		TotalPages:  totalPages,
		CurrentPage: page,
		PerPage:     limit,
		HasNextPage: page < totalPages,
		HasPrevPage: page > 1,
	}

	responseList := make([]dto.StallResponseData, 0, len(stalls))
	for _, s := range stalls {
		responseList = append(responseList, *u.mapToSimpleResponse(&s))
	}

	return responseList, meta, nil
}

func (u *StallUsecase) mapToDetailResponse(s *entity.Stall, owner *entity.User) any {
	// 1. Owner Summary Mapping
	ownerSummary := dto.OwnerProfileSummary{
		ID:          s.StallOwnerID.String(),
		Name:        "Unknown Owner",
		Contact:     "",
		AvatarURL:   "",
		Rating:      5.0,
		ReviewCount: 0,
		JoinedYear:  fmt.Sprintf("%d", time.Now().Year()),
	}

	if owner != nil {
		ownerSummary.Name = owner.Name
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

	// 2. Media Mapping (Type Assertion untuk any ke string)
	facilityImgs := make([]dto.FacilityImageDTO, 0, len(s.DisplayMedia.FacilityImages))
	for idx, imgVal := range s.DisplayMedia.FacilityImages {
		imgStr := ""
		if str, ok := imgVal.(string); ok {
			imgStr = str
		} else if imgMap, ok := imgVal.(map[string]any); ok {
			if urlVal, exists := imgMap["url"].(string); exists {
				imgStr = urlVal
			}
		}

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

	// 3. Address Mapping
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

	// 4. Landmarks Mapping
	nearbyLandmarks := make([]dto.NearbyLandmarkDTO, 0, len(s.NearbyLandmarks))
	for _, l := range s.NearbyLandmarks {
		nearbyLandmarks = append(nearbyLandmarks, dto.NearbyLandmarkDTO{
			CategoryValue: l.CategoryValue,
			Name:          l.Name,
			DistanceKm:    l.DistanceKm,
		})
	}

	// 5. Pricing Mapping
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

	// 6. Polymorphic Mapping berdasarkan PermanenceType
	switch s.PermanenceType {
	case entity.StallPermanenceSemi:
		parentComplex := ""
		if s.ParentComplexName != nil {
			parentComplex = *s.ParentComplexName
		}
		opHours := struct {
			OpeningTime string `json:"openingTime"`
			ClosingTime string `json:"closingTime"`
			Is24Hours   bool   `json:"is24Hours"`
		}{}
		if s.OperatingHours != nil {
			opHours.OpeningTime = s.OperatingHours.OpeningTime
			opHours.ClosingTime = s.OperatingHours.ClosingTime
			opHours.Is24Hours = s.OperatingHours.Is24Hours
		}

		minMonths := 1
		if s.MinimumLeaseMonths != nil {
			minMonths = *s.MinimumLeaseMonths
		}

		return dto.SemiPermanentStallDetail{
			ID:                    s.ID.String(),
			Title:                 s.Title,
			Description:           desc,
			Media:                 mediaDTO,
			PropertyType:          s.PropertyType,
			PropertyTypeValue:     s.PropertyType,
			Placement:             string(s.Placement),
			ElectricityCapacityVA: getIntVal(s.ElectricityCapacityVA),
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

	case entity.StallPermanenceTemporary:
		eventMeta := struct {
			EventName                      string `json:"eventName,omitempty"`
			EventStartDate                 string `json:"eventStartDate"`
			EventEndDate                   string `json:"eventEndDate"`
			RegistrationDeadlineDaysBefore int    `json:"registrationDeadlineDaysBefore"`
			TotalSlots                     int    `json:"totalSlots"`
			AvailableSlots                 int    `json:"availableSlots"`
		}{}

		if s.EventSchedule != nil {
			eventMeta.EventName = s.EventSchedule.EventName
			eventMeta.EventStartDate = s.EventSchedule.StartDate
			eventMeta.EventEndDate = s.EventSchedule.EndDate
			eventMeta.RegistrationDeadlineDaysBefore = s.EventSchedule.RegistrationDeadlineDays
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

		return dto.TemporaryStallDetail{
			ID:                    s.ID.String(),
			Title:                 s.Title,
			Description:           desc,
			Media:                 mediaDTO,
			PropertyType:          s.PropertyType,
			PropertyTypeValue:     s.PropertyType,
			Placement:             string(s.Placement),
			ElectricityCapacityVA: getIntVal(s.ElectricityCapacityVA),
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

	default: // Permanent
		minMonths := 1
		if s.MinimumLeaseMonths != nil {
			minMonths = *s.MinimumLeaseMonths
		}

		return dto.PermanentStallDetail{
			ID:                    s.ID.String(),
			Title:                 s.Title,
			Description:           desc,
			Media:                 mediaDTO,
			PropertyType:          s.PropertyType,
			PropertyTypeValue:     s.PropertyType,
			Placement:             string(s.Placement),
			ElectricityCapacityVA: getIntVal(s.ElectricityCapacityVA),
			Address:               addressDTO,
			NearbyLandmarks:       nearbyLandmarks,
			Pricing:               pricingDTO,
			FacilityValues:        s.FacilityValues,
			HouseRules:            s.HouseRules,
			Rating:                s.RatingAvg,
			ReviewCount:           s.ReviewCount,
			Owner:                 ownerSummary,
			PermanenceType:        "permanent",
			SizeSqm:               getFloatVal(s.SizeSqm),
			Dimensions: struct {
				LengthMeters float64 `json:"lengthMeters"`
				WidthMeters  float64 `json:"widthMeters"`
			}{
				LengthMeters: getFloatVal(s.LengthMeters),
				WidthMeters:  getFloatVal(s.WidthMeters),
			},
			FloorLevel: getIntVal(s.FloorLevel),
			LeaseRules: dto.PermanentLeaseRules{
				MinimumLeaseMonths: minMonths,
				StartDateOptions:   s.StartDateOptions,
				UtilityTerms:       utilityTerms,
			},
		}
	}
}

func (u *StallUsecase) mapToSimpleResponse(s *entity.Stall) *dto.StallResponseData {
	var opHours *dto.OperatingHoursDTO
	if s.OperatingHours != nil {
		opHours = &dto.OperatingHoursDTO{
			OpeningTime: s.OperatingHours.OpeningTime,
			ClosingTime: s.OperatingHours.ClosingTime,
			Is24Hours:   s.OperatingHours.Is24Hours,
		}
	}

	var eventSched *dto.EventScheduleDTO
	if s.EventSchedule != nil {
		eventSched = &dto.EventScheduleDTO{
			EventName:                s.EventSchedule.EventName,
			StartDate:                s.EventSchedule.StartDate,
			EndDate:                  s.EventSchedule.EndDate,
			RegistrationDeadlineDays: s.EventSchedule.RegistrationDeadlineDays,
		}
	}

	var slotInf *dto.SlotInfoDTO
	if s.SlotInfo != nil {
		slotInf = &dto.SlotInfoDTO{
			TotalSlots:     s.SlotInfo.TotalSlots,
			AvailableSlots: s.SlotInfo.AvailableSlots,
		}
	}

	var opDays, attReq, cancelPol *string
	if s.EventOperatingDays != nil {
		v := string(*s.EventOperatingDays)
		opDays = &v
	}
	if s.EventAttendanceReq != nil {
		v := string(*s.EventAttendanceReq)
		attReq = &v
	}
	if s.EventCancellationPolicy != nil {
		v := string(*s.EventCancellationPolicy)
		cancelPol = &v
	}

	facilityImgs := make([]string, 0, len(s.DisplayMedia.FacilityImages))
	for _, imgVal := range s.DisplayMedia.FacilityImages {
		if str, ok := imgVal.(string); ok {
			facilityImgs = append(facilityImgs, str)
		} else if imgMap, ok := imgVal.(map[string]any); ok {
			if urlVal, exists := imgMap["url"].(string); exists {
				facilityImgs = append(facilityImgs, urlVal)
			}
		}
	}

	landmarks := make([]string, 0, len(s.NearbyLandmarks))
	for _, l := range s.NearbyLandmarks {
		landmarks = append(landmarks, l.Name)
	}

	return &dto.StallResponseData{
		ID:                      s.ID.String(),
		StallOwnerID:            s.StallOwnerID.String(),
		Title:                   s.Title,
		Description:             s.Description,
		PropertyType:            s.PropertyType,
		PermanenceType:          string(s.PermanenceType),
		Placement:               string(s.Placement),
		SizeSqm:                 s.SizeSqm,
		LengthMeters:            s.LengthMeters,
		WidthMeters:             s.WidthMeters,
		FloorLevel:              s.FloorLevel,
		ElectricityCapacityVA:   s.ElectricityCapacityVA,
		ParentComplexName:       s.ParentComplexName,
		OperatingHours:          opHours,
		EventSchedule:           eventSched,
		SlotInfo:                slotInf,
		StreetAddress:           s.StreetAddress,
		Suburb:                  s.Suburb,
		District:                s.District,
		City:                    s.City,
		Province:                s.Province,
		Country:                 s.Country,
		CountryCode:             s.CountryCode,
		PostalCode:              s.PostalCode,
		Latitude:                s.Latitude,
		Longitude:               s.Longitude,
		MapURL:                  s.MapURL,
		EmbeddedMapURL:          s.EmbeddedMapURL,
		NearbyLandmarks:         landmarks,
		AllowedPaymentCycles:    s.AllowedPaymentCycles,
		DailyRate:               s.DailyRate,
		MonthlyRate:             s.MonthlyRate,
		QuarterlyRate:           s.QuarterlyRate,
		SemesterlyRate:          s.SemesterlyRate,
		YearlyRate:              s.YearlyRate,
		SecurityDeposit:         s.SecurityDeposit,
		MinimumLeaseMonths:      s.MinimumLeaseMonths,
		MinimumLeaseDays:        s.MinimumLeaseDays,
		StartDateOptions:        s.StartDateOptions,
		EventOperatingDays:      opDays,
		EventAttendanceReq:      attReq,
		EventCancellationPolicy: cancelPol,
		UtilityTerms:            s.UtilityTerms,
		FacilityValues:          s.FacilityValues,
		AllowedBusinessTypeIDs:  s.AllowedBusinessTypeIDs,
		HouseRules:              s.HouseRules,
		DisplayMedia: dto.DisplayMediaDTO{
			MainImage:      s.DisplayMedia.MainImage,
			FacilityImages: facilityImgs,
		},
		LegalDocuments: s.LegalDocuments,
		RatingAvg:      s.RatingAvg,
		ReviewCount:    s.ReviewCount,
		IsPublished:    s.IsPublished,
		CreatedAt:      s.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      s.UpdatedAt.Format(time.RFC3339),
	}
}

func getIntVal(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func getFloatVal(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}
