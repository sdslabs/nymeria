// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sdslabs/nymeria/internal/config"
	"github.com/sdslabs/nymeria/internal/database/schema"
)

// The global database instance
var DB *gorm.DB

// connect establishes a connection to the PostgreSQL database using GORM
func connect() (*gorm.DB, error) {
	if config.AppConfig.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASS environment variable is required")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DBHost, config.AppConfig.DBUser, config.AppConfig.DBPassword, config.AppConfig.DBName, config.AppConfig.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Enable uuid-ossp extension for uuid_generate_v4() function
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return nil, fmt.Errorf("failed to create uuid-ossp extension: %w", err)
	}

	if err := db.AutoMigrate(&schema.User{}, &schema.Organization{}, &schema.Application{}, &schema.OTP{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return db, nil
}

// Init initializes the database connection
func Init() error {
	db, err := connect()
	if err != nil {
		return fmt.Errorf("database initialization failed: %w", err)
	}

	DB = db
	log.Println("Database connected successfully")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB == nil {
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	return sqlDB.Close()
}
