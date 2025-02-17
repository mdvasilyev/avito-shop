package db

import (
	"database/sql"
	"log/slog"

	"github.com/pressly/goose/v3"
)

func RunMigrations(lgr *slog.Logger, db *sql.DB, migrationsPath string) error {
	lgr.Info("Running migrations")

	err := goose.Up(db, migrationsPath)
	if err != nil {
		lgr.Error("Error while goose up operation", "error", err)
		return err
	}

	return nil
}
