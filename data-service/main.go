package main

import (
	"context"
	"os"
	"time"

	"github.com/bektosh03/test-crud/common/environment"
	"github.com/bektosh03/test-crud/common/server"
	"github.com/bektosh03/test-crud/data-service/adapters"
	"github.com/bektosh03/test-crud/data-service/app"
	"github.com/bektosh03/test-crud/data-service/config"
	"github.com/bektosh03/test-crud/common/postgres"
	"github.com/bektosh03/test-crud/data-service/ports"
	"github.com/bektosh03/test-crud/genprotos/datapb"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	db, err := postgres.Connect(ctx, postgres.Config{ConnString: cfg.PostgresConnString()})
	if err != nil {
		logrus.Panicf("failed to connect to db: %v; connString: %s", err, cfg.PostgresConnString())
	}
	defer db.Close()

	app := app.New(adapters.NewPostgresRepository(db))
	server.RunGRPCServer(cfg.ListenAddress(), func(s *grpc.Server) {
		svc := ports.NewGrpcServer(&app)
		datapb.RegisterDataServiceServer(s, svc)
	})
}
