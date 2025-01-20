package main

import (
	"entdemo/config"
	"entdemo/internal/handler"
	"entdemo/internal/repository"
	"entdemo/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	runServer()
}

func runServer() {
	cfg := config.LoadEnvVariables()

	db, err := repository.NewDb(cfg.ConnectionString)
	if err != nil {
		logrus.Fatal(err)
	}
	hdl := handler.NewHandler(db)

	srv, err := server.NewServer(hdl)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithField("cfg", cfg).Info("Starting server with these configuration variables")
	err = srv.Run(cfg.Port)
	if err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	}
}
