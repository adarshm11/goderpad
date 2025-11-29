package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	const envFile = ".env"
	if _, err := os.Stat(envFile); err == nil {
		if err := godotenv.Load(envFile); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	} else if os.IsNotExist(err) {
		log.Printf("Warning: .env file not found; environment variables must be set via system")
	} else {
		log.Fatalf("Error checking .env file: %v", err)
	}
}

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return value, nil
}
