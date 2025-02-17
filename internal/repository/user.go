package repository

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/mdvasilyev/avito-shop/internal/model"
)

type UserRepository struct {
	lgr *slog.Logger
	db  *sql.DB
}

func NewUserRepository(lgr *slog.Logger, db *sql.DB) *UserRepository {
	return &UserRepository{lgr: lgr, db: db}
}

func (rps *UserRepository) BeginTx() (*sql.Tx, error) {
	rps.lgr.Info("Beginning transaction")

	return rps.db.Begin()
}

func (rps *UserRepository) GetUserCoins(userID int) (int, error) {
	rps.lgr.Info("Getting user coins")

	var coins int
	query := "SELECT coins FROM employee WHERE id = $1"

	err := rps.db.QueryRow(query, userID).Scan(&coins)
	if err != nil {
		rps.lgr.Error("Error while scanning coins", "error", err)
		return 0, errors.New("error while scanning coins")
	}

	return coins, nil
}

func (rps *UserRepository) GetMerchNameByID(id int) (string, error) {
	rps.lgr.Info("Getting merch name by id")

	row := rps.db.QueryRow(
		"SELECT name FROM merch WHERE id = $1",
		id,
	)

	var merchName string

	if err := row.Scan(&merchName); err != nil {
		rps.lgr.Error("Merch name not found", "error", err)
		return "", errors.New("Merch name not found")
	}

	return merchName, nil
}

func (rps *UserRepository) GetUserInventory(userID int) ([]model.InventoryItem, error) {
	rps.lgr.Info("Getting user inventory")

	rows, err := rps.db.Query(
		"SELECT merch_id, quantity FROM purchase WHERE employee_id = $1",
		userID,
	)
	if err != nil {
		rps.lgr.Error("Error while getting merch", "error", err)
		return nil, errors.New("error while getting merch")
	}
	defer rows.Close()

	inventory := make([]model.InventoryItem, 0)
	for rows.Next() {
		var item model.InventoryItem
		var merchID int
		var merchName string

		if err = rows.Scan(merchID, &item.Quantity); err != nil {
			rps.lgr.Error("Error while scanning inventory item", "error", err)
			return nil, errors.New("error while scanning inventory item")
		}

		merchName, err = rps.GetMerchNameByID(merchID)
		if err != nil {
			rps.lgr.Error("Error while getting merch name by id", "error", err)
			return nil, errors.New("error while getting merch name by id")
		}
		item.Name = merchName
		inventory = append(inventory, item)
	}

	return inventory, nil
}

func (rps *UserRepository) GetCoinTransactions(userID int) (model.CoinHistory, error) {
	rps.lgr.Info("Getting coin transactions")

	coinHistory := model.CoinHistory{Received: make([]model.CoinTransaction, 0), Sent: make([]model.CoinTransaction, 0)}

	rowsSent, err := rps.db.Query(
		"SELECT receiver_id, quantity FROM transactions WHERE sender_id = $1",
		userID,
	)
	if err != nil {
		rps.lgr.Error("Error while querying db", "error", err)
		return coinHistory, errors.New("error while querying db")
	}
	defer rowsSent.Close()

	for rowsSent.Next() {
		var coinTransaction model.CoinTransaction
		if err = rowsSent.Scan(&coinTransaction.ReceiverID, &coinTransaction.Quantity); err != nil {
			rps.lgr.Error("Error while scanning transaction", "error", err)
			return coinHistory, errors.New("error while scanning transaction")
		}
		coinHistory.Sent = append(coinHistory.Sent, coinTransaction)
	}

	rowsReceived, err := rps.db.Query(
		"SELECT sender_id, quantity FROM transactions WHERE receiver_id = $1",
		userID,
	)
	if err != nil {
		rps.lgr.Error("Error while querying db", "error", err)
		return coinHistory, errors.New("error while querying db")
	}
	defer rowsReceived.Close()

	for rowsReceived.Next() {
		var coinTransaction model.CoinTransaction
		if err = rowsReceived.Scan(&coinTransaction.SenderID, &coinTransaction.Quantity); err != nil {
			rps.lgr.Error("Error while scanning transaction", "error", err)
			return coinHistory, errors.New("error while scanning transaction")
		}
		coinHistory.Received = append(coinHistory.Received, coinTransaction)
	}

	return coinHistory, nil
}
