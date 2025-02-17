package repository

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/mdvasilyev/avito-shop/internal/model"
)

type CoinRepository struct {
	lgr *slog.Logger
	db  *sql.DB
}

func NewCoinRepository(lgr *slog.Logger, db *sql.DB) *CoinRepository {
	return &CoinRepository{lgr: lgr, db: db}
}

func (rps *CoinRepository) BeginTx() (*sql.Tx, error) {
	rps.lgr.Info("Beginning transaction")

	return rps.db.Begin()
}

func (rps *CoinRepository) GetUserForUpdate(userID int) (model.CoinUser, error) {
	rps.lgr.Info("Getting user coin")

	row := rps.db.QueryRow(
		"SELECT id, coins FROM employee WHERE id = $1 FOR UPDATE",
		userID,
	)

	var user model.CoinUser

	if err := row.Scan(&user.ID, &user.Coins); err != nil {
		rps.lgr.Error("Error while scanning user", "error", err)
		return model.CoinUser{}, errors.New("error while scanning user")
	}

	return user, nil
}

func (rps *CoinRepository) UpdateBalance(userID, coins int) error {
	rps.lgr.Info("Updating balance")

	_, err := rps.db.Exec(
		"UPDATE employee SET coins = $1 WHERE id = $2",
		coins,
		userID,
	)

	return err
}

func (rps *CoinRepository) CreateTransaction(transaction *model.CoinTransaction) error {
	rps.lgr.Info("Creating transaction")

	_, err := rps.db.Exec(
		"INSERT INTO transactions (sender_id, receiver_id, quantity) VALUES ($1, $2, $3)",
		transaction.SenderID,
		transaction.ReceiverID,
		transaction.Quantity,
	)

	return err
}
