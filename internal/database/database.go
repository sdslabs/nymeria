// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// The global database instance
var DB *gorm.DB

// loadConfig loads database configuration from environment variables
func loadConfig() *Config {
	return &Config{
		// TODO: update this
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		User:     getEnvOrDefault("DB_USER", "nymeria"),
		Password: getEnvOrDefault("DB_PASS", "password"),
		DBName:   getEnvOrDefault("DB_NAME", "nymeria"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
	}
}

// getEnvOrDefault retrieves the value of the environment variable named by key.
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// connect establishes a connection to the PostgreSQL database using GORM
func connect(config *Config) (*gorm.DB, error) {
	if config.Password == "" {
		return nil, fmt.Errorf("DB_PASS environment variable is required")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&User{}, &Organization{}, &Application{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return db, nil
}

// Init initializes the database connection
func Init() error {
	config := loadConfig()

	db, err := connect(config)
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
