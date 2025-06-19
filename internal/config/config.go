// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package config

import (
	"os"
	"strconv"
)

// Global configuration instance
var AppConfig *Config

// GetEnvOrDefault retrieves the value of the environment variable named by key.
// If the environment variable is not set or is empty, it returns the default value.
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ParseIntGetEnvOrDefault(key, defaultValue string) int {
	value, err := strconv.Atoi(GetEnvOrDefault(key, defaultValue))
	if err != nil {
		value, _ = strconv.Atoi(defaultValue)
	}
	return value
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		DBHost:     GetEnvOrDefault("DB_HOST", "localhost"),
		DBUser:     GetEnvOrDefault("DB_USER", "nymeria"),
		DBPassword: GetEnvOrDefault("DB_PASS", "password"),
		DBName:     GetEnvOrDefault("DB_NAME", "nymeria"),
		DBPort:     GetEnvOrDefault("DB_PORT", "5432"),

		CSRFSecret: GetEnvOrDefault("CSRF_SECRET", "csrf-secret-key-32"),
		CSRFMaxAge: ParseIntGetEnvOrDefault("CSRF_MAX_AGE", "2"),
		JWTSecret:  GetEnvOrDefault("JWT_SECRET", "jwt-secret-key-32"),
		JWTMaxAge:  ParseIntGetEnvOrDefault("JWT_MAX_AGE", "2"),

		EnvMode: GetEnvOrDefault("ENV_MODE", "development"),
	}
}

// Init initializes the global configuration
func Init() {
	AppConfig = LoadConfig()
}
