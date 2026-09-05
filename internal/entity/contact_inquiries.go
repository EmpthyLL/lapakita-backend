package entity

import (
	"time"

	"github.com/google/uuid"
)

type ContactInquiry struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Email       string    `gorm:"type:varchar(255);not null" json:"email"`
	Whatsapp    *string   `gorm:"type:varchar(32)" json:"whatsapp"`
	Persona     string    `gorm:"type:varchar(64);not null" json:"persona"`
	InquiryType string    `gorm:"type:varchar(64);not null" json:"inquiry_type"`
	Message     string    `gorm:"type:text;not null" json:"message"`
	Status      string    `gorm:"type:varchar(32);default:'unread'" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (ContactInquiry) TableName() string {
	return "contact_inquiries"
}
