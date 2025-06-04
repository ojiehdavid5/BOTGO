package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

// LoadConfig loads the configuration from the .env file and returns the value of the specified key.
func LoadConfig(key string) (string, error) {
    err := godotenv.Load()
    if err != nil {
        return "", fmt.Errorf("error loading .env file: %w", err)
    }

    value, exists := os.LookupEnv(key)
    if !exists {
        return "", fmt.Errorf("environment variable %s does not exist", key)
    }

    return value, nil
}
