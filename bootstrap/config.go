package bootstrap

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`

	Debug      bool   `envconfig:"DATABASE_DEBUG" default:"false"`
	Migrate    bool   `envconfig:"DATABASE_MIGRATE" default:"false"`
	ServerAddr string `envconfig:"SERVER_ADDR" default:":8002"`
	PageSize   int    `envconfig:"DEFAULT_PAGE_SIZE" default:"10"`
}

func LoadConfig() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	return &cfg
}
