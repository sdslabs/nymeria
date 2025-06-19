package schema

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string         `gorm:"not null;unique"`
	Users        []*User        `gorm:"many2many:user_organizations;"`
	Applications []*Application `gorm:"many2many:application_organizations;"`
	UserRoles    []UserRole     `gorm:"foreignKey:OrganizationID"`
	CreatedAt    time.Time      `gorm:"not null"`
}
