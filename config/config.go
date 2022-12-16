package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SlackBaseURL   string `envconfig:"SLACK_BASE_URL" default:""`
	SlackAPIKey    string `envconfig:"SLACK_API_KEY" default:""`
	SlackBotAPIKey string `envconfig:"SLACK_BOT_API_KEY" default:""`
	HTTPPort       string `envconfig:"HTTP_PORT" default:"8080"`
	Database       Database
	RedisHostName  string `envconfig:"REDIS_HOST_NAME" default:"localhost"`
	RedisUserName  string `envconfig:"REDIS_USER_NAME" default:""`
	RedisPassword  string `envconfig:"REDIS_PASSWORD" default:""`
	RedisPort      int    `envconfig:"REDIS_PORT" default:"6379"`
	ScheduleCron   ScheduleCron
}

type Database struct {
	DBHostName string `envconfig:"DB_HOST_NAME" default:""`
	DBUserName string `envconfig:"DB_USER_NAME" default:""`
	DBPassword string `envconfig:"DB_PASSWORD" default:""`
	DBPort     string `envconfig:"DB_PORT" default:""`
	DBName     string `envconfig:"DB_NAME" default:""`
}

type ScheduleCron struct {
	UpdateUsersGroupMinutely string `envconfig:"UPDATE_USERS_GROUP_MINUTELY_CRON" default:"* * * * *"`
	UpdateUsersGroupHourly   string `envconfig:"UPDATE_USERS_GROUP_HOURLY_CRON" default:"0 * * * *"`
	UpdateUsersGroupDaily    string `envconfig:"UPDATE_USERS_GROUP_DAILY_CRON" default:"* 9 * * *"`
	UpdateUsersGroupWeekly   string `envconfig:"UPDATE_USERS_GROUP_WEEKLY_CRON" default:"* 9 * * 1"`
	UpdateUsersGroupMonthly  string `envconfig:"UPDATE_USERS_GROUP_MONTHLY_CRON" default:"* 9 1 * *"`
}

func Get() Config {
	_ = godotenv.Overload()
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
