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

	err = environment.Set(cfg.Environment)
	if err != nil {
		return Config{}, err
	}

	return Config{
		host:            cfg.Host,
		port:            cfg.Port,
		dataServiceHost: cfg.DataServiceHost,
		dataServicePort: cfg.DataServicePort,
		postServiceHost: cfg.PostServiceHost,
		postServicePort: cfg.PostServicePort,
	}, nil
}

type Config struct {
	host            string
	port            int
	dataServiceHost string
	dataServicePort int
	postServiceHost string
	postServicePort int
}

func (c Config) ListenAddress() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c Config) DataServiceAddr() string {
	return fmt.Sprintf("%s:%d", c.dataServiceHost, c.dataServicePort)
}

func (c Config) PostServiceAddr() string {
	return fmt.Sprintf("%s:%d", c.postServiceHost, c.postServicePort)
}

type config struct {
	Environment     string `envconfig:"ENVIRONMENT" default:"development"`
	Host            string `envconfig:"HOST" default:"localhost"`
	Port            int    `envconfig:"PORT" default:"8000"`
	DataServiceHost string `envconfig:"DATA_SERVICE_HOST" default:"localhost"`
	DataServicePort int    `envconfig:"DATA_SERVICE_PORT" default:"8000"`
	PostServiceHost string `envconfig:"POST_SERVICE_HOST" default:"localhost"`
	PostServicePort int    `envconfig:"POST_SERVICE_PORT" default:"8001"`
}
