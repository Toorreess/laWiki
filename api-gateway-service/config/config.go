package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ProjectID string
}

func ReadConfig() *Config {
	c := Config{}
	path, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/config/.env", path))
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	c.ProjectID = os.Getenv("PROJECT_ID")

	return &c
}
