package main

import (
	"context"
	"os"
	"time"

	"github.com/bektosh03/test-crud/api-gateway/config"
	"github.com/bektosh03/test-crud/common/environment"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		logrus.Panicf("failed to load config: %v", err)
	}
	if environment.Current() == environment.Development {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"env":               environment.Current().String(),
		"listen addr":       cfg.ListenAddress(),
		"data service addr": cfg.DataServiceAddr(),
		"post service addr": cfg.PostServiceAddr(),
	}).Info("loaded config")
}
