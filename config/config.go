package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// AppConfig holds the application configuration values
type AppConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	ServerPort string
}

// Config is a global variable holding the application configuration
var Config AppConfig

// LoadConfig loads the environment variables from .env or from the system environment
func LoadConfig() {
	// Load .env file if in local development mode
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Read the environment variables
	Config.DBUser = getEnv("DB_USER", "")
	Config.DBPassword = getEnv("DB_PASSWORD", "")
	Config.DBName = getEnv("DB_NAME", "")
	Config.DBHost = getEnv("DB_HOST", "")
	Config.DBPort = getEnv("DB_PORT", "")
	Config.ServerPort = getEnv("SERVER_PORT", "")

	// Ensure required environment variables are set
	if Config.DBUser == "" || Config.DBName == "" {
		log.Fatal("Missing required environment variables")
	}
}

// getEnv is a helper function to fetch environment variables
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// DatabaseDSN returns the Data Source Name (DSN) to connect to the database
func (cfg *AppConfig) DatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
}
