package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CMSLegalDocument struct {
	ID           uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DocType      string          `gorm:"type:varchar(32);not null;uniqueIndex:idx_doc_type_lang" json:"doc_type"`
	Lang         string          `gorm:"type:varchar(8);not null;default:'en';uniqueIndex:idx_doc_type_lang" json:"lang"`
	Title        string          `gorm:"type:varchar(255);not null" json:"title"`
	Intro        string          `gorm:"type:text;not null" json:"intro"`
	SectionsJSON json.RawMessage `gorm:"type:jsonb;not null" json:"sections_json"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

func (CMSLegalDocument) TableName() string {
	return "cms_legal_documents"
}
