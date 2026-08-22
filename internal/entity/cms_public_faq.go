package entity

import (
	"time"

	"github.com/google/uuid"
)

type CMSPublicFAQ struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Lang          string    `gorm:"type:varchar(8);not null;default:'en'" json:"lang"`
	CategoryID    string    `gorm:"type:varchar(32);not null" json:"category_id"`
	SubTopicTitle string    `gorm:"type:varchar(255);not null" json:"sub_topic_title"`
	Question      string    `gorm:"type:text;not null" json:"question"`
	Answer        string    `gorm:"type:text;not null" json:"answer"`
	RoleType      string    `gorm:"type:varchar(32)" json:"role_type"`
	SortOrder     int       `gorm:"default:0" json:"sort_order"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (CMSPublicFAQ) TableName() string {
	return "cms_public_faqs"
}
