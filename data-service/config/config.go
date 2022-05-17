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
		host: cfg.Host,
		port: cfg.Port,
	}, nil
}

type Config struct {
	host string
	port int
}

func (c Config) ListenAddress() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

type config struct {
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
	Host        string `envconfig:"HOST" default:"localhost"`
	Port        int    `envconfig:"PORT" default:"8002"`
}
