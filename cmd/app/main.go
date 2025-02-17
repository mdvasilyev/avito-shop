package main

import (
	"github.com/mdvasilyev/avito-shop/internal/app"
	"github.com/mdvasilyev/avito-shop/internal/config"
	"github.com/mdvasilyev/avito-shop/internal/db"
	"log/slog"
	"os"
)

func main() {
	lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	cfg := config.GetConfig(lgr)

	dbConn, err := db.Connect(lgr, cfg)
	if err != nil {
		lgr.Error("Error while connecting to database")
	}
	defer db.Close(lgr, dbConn)

	err = db.RunMigrations(lgr, dbConn, cfg.MigrationsPath)
	if err != nil {
		lgr.Error("Error while running migrations")
	}

	app.RunServer(lgr, cfg, dbConn)
}
