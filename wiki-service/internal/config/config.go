package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		User     string
		Password string
		DBType   string
		Addr     string
		DBName   string
	}
	Server struct {
		Address string
	}
	ProjectID string
}

func ReadConfig() *Config {
	c := Config{}
	path, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/../.env", path))
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBType = os.Getenv("DB_TYPE")
	c.Database.Addr = os.Getenv("DB_ADDR")
	c.Database.DBName = os.Getenv("DB_NAME")

	c.Server.Address = os.Getenv("SERVER_ADDRESS")

	c.ProjectID = os.Getenv("PROJECT_ID")

	return &c
}
