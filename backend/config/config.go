package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	godotenv.Load()
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return value, nil
}
