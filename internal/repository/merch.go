package repository

import (
	"database/sql"
	"errors"
	"github.com/mdvasilyev/avito-shop/internal/model"
	"log/slog"
)

type MerchRepository struct {
	lgr *slog.Logger
	db  *sql.DB
}

func NewMerchRepository(lgr *slog.Logger, db *sql.DB) *MerchRepository {
	return &MerchRepository{lgr: lgr, db: db}
}

func (rps *MerchRepository) BeginTx() (*sql.Tx, error) {
	rps.lgr.Info("Beginning transaction")

	return rps.db.Begin()
}

func (rps *MerchRepository) GetUserById(userID int) (model.User, error) {
	rps.lgr.Info("Getting user by id")

	row := rps.db.QueryRow(
		"SELECT id, coins FROM employee WHERE id = $1",
		userID,
	)

	var user model.User
	if err := row.Scan(&user.ID, &user.Coins); err != nil {
		rps.lgr.Error("Error while scanning user", "error", err)
		return model.User{}, errors.New("error while scanning user")
	}

	return user, nil
}

func (rps *MerchRepository) GetItemByName(name string) (model.Merch, error) {
	rps.lgr.Info("Getting item by name")

	row := rps.db.QueryRow(
		"SELECT id, name, price FROM merch WHERE name = $1",
		name,
	)

	var merch model.Merch
	if err := row.Scan(&merch.ID, &merch.Name, &merch.Price); err != nil {
		rps.lgr.Error("Error while scanning merch", "error", err)
		return model.Merch{}, errors.New("error while scanning merch")
	}

	return merch, nil
}

func (rps *MerchRepository) UpdateUserBalance(userID, newCoins int) error {
	rps.lgr.Info("Updating user balance")

	_, err := rps.db.Exec(
		"UPDATE employee SET coins = $1 WHERE id = $2",
		newCoins,
		userID,
	)

	return err
}

func (rps *MerchRepository) UpdateInventory(userID, merchID int) error {
	rps.lgr.Info("Updating user inventory")

	_, err := rps.db.Exec(
		`
		INSERT INTO purchase (user_id, merch_id, quantity) 
		VALUES ($1, $2, 1)
		ON CONFLICT (user_id, merch_id) DO UPDATE
		SET quantity = purchase.quantity + 1`,
		userID,
		merchID,
	)

	return err
}
