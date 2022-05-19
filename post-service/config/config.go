package config

import (
	"fmt"

	"github.com/bektosh03/test-crud/common/environment"
	_ "github.com/joho/godotenv/autoload" // automatically load environment variables from .env file
	"github.com/kelseyhightower/envconfig"
)

func Load() (Config, error) {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, err
	}

	if err = environment.Set(cfg.Environment); err != nil {
		return Config{}, err
	}

	return Config{
		host:             cfg.Host,
		port:             cfg.Port,
		postgresHost:     cfg.PostgresHost,
		postgresPort:     cfg.PostgresPort,
		postgresUser:     cfg.PostgresUser,
		postgresPassword: cfg.PostgresPassword,
		postgresDB:       cfg.PostgresDB,
	}, nil
}

type Config struct {
	host             string
	port             int
	postgresHost     string
	postgresPort     int
	postgresUser     string
	postgresPassword string
	postgresDB       string
}

func (c Config) ListenAddress() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c Config) PostgresConnString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.postgresHost, c.postgresPort, c.postgresUser, c.postgresPassword, c.postgresDB,
	)
}

type config struct {
	Environment      string `envconfig:"ENVIRONMENT" default:"development"`
	Host             string `envconfig:"HOST" default:"localhost"`
	Port             int    `envconfig:"PORT" default:"8001"`
	PostgresHost     string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresPort     int    `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresUser     string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB       string `envconfig:"POSTGRES_DB" required:"true"`
}
