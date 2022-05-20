package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/bektosh03/test-crud/api-gateway/api"
	"github.com/bektosh03/test-crud/api-gateway/config"
	"github.com/bektosh03/test-crud/api-gateway/ports"
	"github.com/bektosh03/test-crud/api-gateway/server"
	"github.com/bektosh03/test-crud/common/client"
	"github.com/bektosh03/test-crud/common/environment"
	"github.com/go-chi/chi/v5"
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

	dataServiceClient, closeDataService, err := client.NewDataServiceClient(ctx, cfg.DataServiceAddr())
	if err != nil {
		logrus.Panicf("failed to connect to data service: %v", err)
	}
	defer closeDataService()

	postServiceClient, closePostService, err := client.NewPostServiceClient(ctx, cfg.PostServiceAddr())
	if err != nil {
		logrus.Panicf("failed to connect to post service: %v", err)
	}
	defer closePostService()

	server.RunHTTPServer(cfg.ListenAddress(), func(router chi.Router) http.Handler {
		api.InitRoutes(cfg, router, ports.NewHTTPServer(dataServiceClient, postServiceClient))
		return router
	})
}
