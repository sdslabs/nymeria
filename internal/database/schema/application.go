package schema

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID             uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string          `gorm:"not null"`
	ApplicationURL string          `gorm:"not null"`
	AllowedOrigins []string        `gorm:"type:text[]"`
	RedirectURIs   []string        `gorm:"type:text[]"`
	ClientKey      string          `gorm:"not null;unique"`
	ClientSecret   string          `gorm:"not null"`
	Organizations  []*Organization `gorm:"many2many:application_organizations;"`
	CreatedAt      time.Time       `gorm:"not null"`
}
