package application

import (
	"artforintrovert_test/internal/adapters/http"
	"artforintrovert_test/internal/adapters/mongo"
	"artforintrovert_test/internal/config"
	"artforintrovert_test/internal/domain/service"
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
)

var (
	configPath = flag.String("config-path", "config.yml", "What config file to use")
)

func Start(ctx context.Context) {
	logrus.Info("Reading config")
	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		logrus.WithError(err).Fatal("Can't read config configuration")
	}
	logrus.Info(fmt.Sprintf("Config from %+v was loaded", *configPath))

	logrus.Info("Database initialization")
	db, err := mongo.New(cfg)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to connect to database")
	}
	logrus.Info("Database connected successful")

	logrus.Info("Test service initialization")
	testApp, err := service.New(ctx, db)
	if err != nil {
		logrus.WithError(err).Fatal("Service do not started")
	}
	testApp.Refresh(ctx)
	logrus.Info("Test service initialization successful")

	restServer := http.New(ctx, testApp, cfg)
	logrus.Info("Starting http API on http://localhost:" + cfg.Listen.Ports.Main + cfg.Listen.Paths.Base)
	go func() {
		logrus.Fatal(restServer.Start(ctx))
	}()
}

func Stop() {
	logrus.Fatal("App has stopped")
}
