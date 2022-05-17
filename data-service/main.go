package main

import (
	"context"
	"os"
	"time"

	"github.com/bektosh03/test-crud/common/environment"
	"github.com/bektosh03/test-crud/data-service/app"
	"github.com/bektosh03/test-crud/data-service/config"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		logrus.Panicf("failed to load config: %v", err)
	}
	if environment.Current() == environment.Development {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"env":            environment.Current().String(),
		"listen address": cfg.ListenAddress(),
	}).Info("loaded config")

	app := app.New()
	logrus.Infof("FINISHED WITH ERR: %v", app.DownloadPosts(ctx))
}
