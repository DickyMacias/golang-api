package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	TMDBAPIKey    string
	SessionSecret string
	Port          string
	Environment   string
}

func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		TMDBAPIKey:    getEnv("TMDB_API_KEY", ""),
		SessionSecret: getEnv("SESSION_SECRET", "your-secret-key-change-in-production"),
		Port:          getEnv("PORT", "8080"),
		Environment:   getEnv("ENVIRONMENT", "development"),
	}

	// Validate required environment variables
	if config.DatabaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	if config.TMDBAPIKey == "" {
		log.Fatal("TMDB_API_KEY environment variable is required")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}