package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		DBType string
	}
	Server struct {
		Port string
	}
	ProjectID string
}

func ReadConfig() *Config {
	c := Config{}
	path, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/config/.env", path))
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	c.Database.DBType = os.Getenv("DB_TYPE")

	c.Server.Port = os.Getenv("SERVER_PORT")

	c.ProjectID = os.Getenv("PROJECT_ID")

	return &c
}
