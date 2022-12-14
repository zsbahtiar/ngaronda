package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SlackBaseURL string `envconfig:"SLACK_BASE_URL" default:""`
	SlackAPIKey  string `envconfig:"SLACK_API_KEY" default:""`
	HTTPPort     string `envconfig:"HTTP_PORT" default:"8080"`
	Database     Database
}

type Database struct {
	DBHostName string `envconfig:"DB_HOST_NAME" default:""`
	DBUserName string `envconfig:"DB_USER_NAME" default:""`
	DBPassword string `envconfig:"DB_PASSWORD" default:""`
	DBPort     string `envconfig:"DB_PORT" default:""`
	DBName     string `envconfig:"DB_NAME" default:""`
}

func Get() Config {
	_ = godotenv.Overload()
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
