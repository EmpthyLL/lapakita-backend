package dto

import "time"

type LegalDocumentQueryRequest struct {
	Lang string `form:"lang"`
}

type LegalSubsection struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body"`
}

type LegalSection struct {
	ID          string            `json:"id"`
	Number      string            `json:"number"`
	Title       string            `json:"title"`
	Subsections []LegalSubsection `json:"subsections"`
}

type LegalDocumentResponse struct {
	LastUpdatedAt *time.Time     `json:"last_updated_at"`
	Data          []LegalSection `json:"data"`
}
