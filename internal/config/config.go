package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	App      AppConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port    string
	GinMode string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret          string
	ExpirationHours int
}

// AppConfig holds application configuration
type AppConfig struct {
	Name    string
	Version string
}

var GlobalConfig *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	expirationHours, err := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))
	if err != nil {
		expirationHours = 24
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "football_db"),
		},
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "your-secret-key"),
			ExpirationHours: expirationHours,
		},
		App: AppConfig{
			Name:    getEnv("APP_NAME", "Football Management API"),
			Version: getEnv("APP_VERSION", "1.0.0"),
		},
	}

	GlobalConfig = config
	return config
}

// getEnv gets environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
