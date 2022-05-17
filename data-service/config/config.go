package config

import (
	"fmt"

	"github.com/hiloldevs-hiloltest/book-service/config/environment"
	_ "github.com/joho/godotenv/autoload" // automatically load environment variables from .env file
	"github.com/kelseyhightower/envconfig"
)

func Load() (Config, error) {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, err
	}

	env, err := environment.FromString(cfg.Environment)
	if err != nil {
		return Config{}, err
	}

	return Config{
		env:  env,
		host: cfg.Host,
		port: cfg.Port,
	}, nil
}

type Config struct {
	env  environment.Environment
	host string
	port int
}

func (c Config) ListenAddress() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c Config) Environment() environment.Environment {
	return c.env
}

func (c Config) Host() string {
	return c.host
}

func (c Config) Port() int {
	return c.port
}

type config struct {
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
	Host        string `envconfig:"HOST" default:"localhost"`
	Port        int    `envconfig:"PORT" default:"8002"`
}
