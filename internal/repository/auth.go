package repository

import (
	"database/sql"
	"errors"
	"github.com/mdvasilyev/avito-shop/internal/model"
	"log/slog"
)

type AuthRepository struct {
	lgr *slog.Logger
	db  *sql.DB
}

func NewAuthRepository(lgr *slog.Logger, db *sql.DB) *AuthRepository {
	return &AuthRepository{lgr: lgr, db: db}
}

func (rps *AuthRepository) GetUser(username string) (model.User, error) {
	rps.lgr.Info("Getting user coins")

	var user model.User
	err := rps.db.QueryRow(
		"SELECT id, pass_hash FROM employee WHERE username = $1",
		username,
	).Scan(&user.ID, &user.PassHash)
	if err != nil {
		rps.lgr.Error("Error while scanning user", "error", err)
		return model.User{}, errors.New("error while scanning user")
	}

	return user, nil
}

func (rps *AuthRepository) UpdateUser(username, hashedPass string) (int, error) {
	rps.lgr.Info("Getting user coins")

	var userID int
	err := rps.db.QueryRow(
		"INSERT INTO employee (username, pass_hash) VALUES ($1, $2) RETURNING id",
		username, hashedPass,
	).Scan(&userID)
	if err != nil {
		rps.lgr.Error("Error while scanning user", "error", err)
		return 0, errors.New("error while scanning user")
	}

	return userID, nil
}
