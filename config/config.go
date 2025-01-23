package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBSSLMode         string
	GRPCPort          string
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
}

// LoadConfig loads the application configuration from the .env file.
func LoadConfig() *Config {
	// Load the .env file into environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found. Using environment variables.")
	}

	return &Config{
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBUser:            getEnv("DB_USER", "postgres"),
		DBPassword:        getEnv("DB_PASSWORD", ""),
		DBName:            getEnv("DB_NAME", "user_db"),
		DBSSLMode:         getEnv("DB_SSLMODE", "disable"),
		GRPCPort:          getEnv("GRPC_PORT", "50051"),
		Auth0Domain:       getEnv("AUTH0_DOMAIN", ""),
		Auth0ClientID:     getEnv("AUTH0_CLIENT_ID", ""),
		Auth0ClientSecret: getEnv("AUTH0_CLIENT_SECRET", ""),
	}
}

// getEnv retrieves environment variables or provides a default value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
