package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PayoutMethod struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null;index:idx_payout_methods_user_id" json:"user_id"`
	PayoutType        string    `gorm:"type:varchar(32);not null;default:'bank_account'" json:"payout_type"` // 'bank_account' atau 'e_wallet'
	ProviderCode      string    `gorm:"type:varchar(32);not null" json:"provider_code"`                      // e.g. 'BCA', 'MANDIRI', 'GOPAY', 'OVO'
	ProviderName      string    `gorm:"type:varchar(128);not null" json:"provider_name"`                     // e.g. 'Bank Central Asia' atau 'GoPay'
	AccountNumber     string    `gorm:"type:varchar(64);not null" json:"account_number"`                     // Nomor Rekening atau Nomor HP E-Wallet
	AccountHolderName string    `gorm:"type:varchar(255);not null" json:"account_holder_name"`               // Nama Pemilik Rekening/Akun
	IsPrimary         bool      `gorm:"default:false" json:"is_primary"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (PayoutMethod) TableName() string {
	return "payout_methods"
}

func (p *PayoutMethod) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
