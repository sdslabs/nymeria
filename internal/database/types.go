// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package database

import "gorm.io/gorm"

// Config holds database configuration
type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

// User holds user info in the database
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	// Phone    string `json:"phone" gorm:"unique; not null; type:char(10)"`
	// Email    string `json:"email" gorm:"unique; not null; type:varchar(255)"`
	Phone    string `json:"phone" gorm:"not null; type:char(10)"`
	Email    string `json:"email" gorm:"not null; type:varchar(255)"`
	Role     string `json:"role" gorm:"default:user; not null"`
	ImageURL string `json:"image_url" gorm:"default:''; not null; type:varchar(255)"`
	// GitHubID         string `json:"github_id" gorm:"unique; not null; type:varchar(255)"`
	GitHubID         string `json:"github_id" gorm:"not null; type:varchar(255)"`
	AccountStatus    string `json:"account_status" gorm:"default:active; not null"` // active, banned, account takeover, etc.
	Verified         bool   `json:"verified" gorm:"default:false; not null"`
	AdditionalEmails string `json:"additional_emails" gorm:"type:text; default:'{}'"`

	// TODO: add apps after GBM discussion
}

type Organization struct {
	gorm.Model
	Name string `json:"name" gorm:"not null; type:varchar(255)"`
}

type Application struct {
	gorm.Model
	Name           string       `json:"name" gorm:"not null; type:varchar(255)"`
	Description    string       `json:"description" gorm:"not null; type:varchar(255)"`
	AppURL         string       `json:"app_url" gorm:"not null; type:varchar(255)"`
	AllowedDomains string       `json:"allowed_domains" gorm:"not null; type:varchar(255)"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"organization" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OrganizationID"`
}
