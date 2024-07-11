package config

import (
	"os"

	"github.com/joho/godotenv"
)

// this function will load the .env file if the GO_ENV environment variable is not set
func LoadENV() error {
	if goEnv := os.Getenv("GO_ENV"); goEnv == "dev" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
