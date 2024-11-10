package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port        string
		WikiPort    string
		EntryPort   string
		CommentPort string
	}

	ProjectID string `yml:"project_id"`
}

func ReadConfig() *Config {
	c := Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "api-gateway-service", "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading envvars: %s", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("Error unmarshalling envvars: %v", err)
	}

	os.Setenv("PROJECT_ID", c.ProjectID)

	return &c
}