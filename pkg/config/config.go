package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnvironment(key string, file string) string {
	if loadEnvironment(file) {
		getEnvValue := os.Getenv(key)
		return getEnvValue
	}
	return fmt.Sprint("Environment file loading error")
}

func loadEnvironment(filename string) bool {
	envLoadError := godotenv.Load(filename)
	if envLoadError != nil {
		return false
	}
	return true
}
