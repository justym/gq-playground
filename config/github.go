package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetGithubToken() (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return "", err
	}
	result := os.Getenv("GITHUB_TOKEN")
	return result, nil
}
