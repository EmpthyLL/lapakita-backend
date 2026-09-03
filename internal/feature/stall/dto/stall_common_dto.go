package dto

import "encoding/json"

type LandmarkRadiusEntry struct {
	Landmark string `json:"landmark" form:"landmark"`
	Radius   string `json:"radius" form:"radius"`
}

type OperatingHoursDTO struct {
	OpeningTime string `json:"opening_time" binding:"required"`
	ClosingTime string `json:"closing_time" binding:"required"`
	Is24Hours   bool   `json:"is_24_hours"`
}

type EventScheduleDTO struct {
	EventName                string `json:"event_name" binding:"required"`
	StartDate                string `json:"start_date" binding:"required"`
	EndDate                  string `json:"end_date" binding:"required"`
	RegistrationDeadlineDays int    `json:"registration_deadline_days"`
}

type SlotInfoDTO struct {
	TotalSlots     int `json:"total_slots" binding:"required"`
	AvailableSlots int `json:"available_slots"`
}

type DisplayMediaDTO struct {
	MainImage      string   `json:"main_image" binding:"required"`
	FacilityImages []string `json:"facility_images"`
}

type CompactStallResponse struct {
	*PermanentStallResponse     `json:",inline"`
	*SemiPermanentStallResponse `json:",inline"`
	*TemporaryStallResponse     `json:",inline"`
}

func (r CompactStallResponse) MarshalJSON() ([]byte, error) {
	switch {
	case r.PermanentStallResponse != nil:
		return json.Marshal(r.PermanentStallResponse)
	case r.SemiPermanentStallResponse != nil:
		return json.Marshal(r.SemiPermanentStallResponse)
	case r.TemporaryStallResponse != nil:
		return json.Marshal(r.TemporaryStallResponse)
	default:
		return []byte("null"), nil
	}
}

type StallDetailResponse struct {
	*PermanentStallDetail     `json:",inline"`
	*SemiPermanentStallDetail `json:",inline"`
	*TemporaryStallDetail     `json:",inline"`
}

func (r StallDetailResponse) MarshalJSON() ([]byte, error) {
	switch {
	case r.PermanentStallDetail != nil:
		return json.Marshal(r.PermanentStallDetail)
	case r.SemiPermanentStallDetail != nil:
		return json.Marshal(r.SemiPermanentStallDetail)
	case r.TemporaryStallDetail != nil:
		return json.Marshal(r.TemporaryStallDetail)
	default:
		return []byte("null"), nil
	}
}
