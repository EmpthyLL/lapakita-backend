package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankAccount struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	BankCode          string    `gorm:"type:varchar(32);not null" json:"bank_code"`
	BankName          string    `gorm:"type:varchar(128);not null" json:"bank_name"`
	AccountNumber     string    `gorm:"type:varchar(64);not null" json:"account_number"`
	AccountHolderName string    `gorm:"type:varchar(255);not null" json:"account_holder_name"`
	IsPrimary         bool      `gorm:"default:false" json:"is_primary"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (BankAccount) TableName() string {
	return "bank_accounts"
}

func (b *BankAccount) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
