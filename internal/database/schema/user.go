package schema

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	USER       Role = "user"
	ADMIN      Role = "admin"
	SUPERADMIN Role = "superadmin"
)

type UserRole struct {
	ID             uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID     `gorm:"type:uuid;index"`
	OrganizationID uuid.UUID     `gorm:"type:uuid;index"`
	Role           Role          `gorm:"not null"`
	User           *User         `gorm:"foreignKey:UserID"`
	Organization   *Organization `gorm:"foreignKey:OrganizationID"`
}

type User struct {
	ID               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email            string          `gorm:"not null;unique"`
	Password         string          `gorm:"not null"`
	Phone            string          `gorm:"not null"`
	Username         string          `gorm:"not null"`
	ImageURL         string          `gorm:"not null"`
	GitHubID         string          `gorm:"not null"`
	AccountStatus    string          `gorm:"not null"`
	Verified         bool            `gorm:"not null"`
	AdditionalEmails string          `gorm:"not null"`
	Organizations    []*Organization `gorm:"many2many:user_organizations;"`
	Roles            []UserRole      `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time       `gorm:"not null"`
}
