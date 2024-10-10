package database

import (
	"time"

	"github.com/google/uuid"
)

type Otps struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	OTP       string    `gorm:"size:6"`
	OTPExpiry time.Time `gorm:""`
	IsActive  bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy string    `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy string    `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;default:now()"`
}