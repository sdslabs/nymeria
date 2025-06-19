package schema

import (
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email     string    `gorm:"not null"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Verified  bool      `gorm:"not null"`
}
