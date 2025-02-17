package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mdvasilyev/avito-shop/internal/config"
	"log/slog"
)

func Connect(lgr *slog.Logger, config *config.Config) (*sql.DB, error) {
	lgr.Info("Connecting to database")

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		lgr.Error("Error while connecting to database", "error", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		lgr.Error("Error while pinging database", "error", err)
		return nil, err
	}

	lgr.Info("Connecting to database is finished")
	return db, nil
}

func Close(lgr *slog.Logger, db *sql.DB) {
	lgr.Info("Closing database connection")

	err := db.Close()
	if err != nil {
		lgr.Error("Error while closing database", "error", err)
	}
}
