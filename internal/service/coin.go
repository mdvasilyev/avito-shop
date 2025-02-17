package service

import (
	"errors"
	"log/slog"

	"github.com/mdvasilyev/avito-shop/internal/model"
	"github.com/mdvasilyev/avito-shop/internal/repository"
)

type CoinService struct {
	lgr  *slog.Logger
	repo *repository.CoinRepository
}

func NewCoinService(lgr *slog.Logger, repo *repository.CoinRepository) *CoinService {
	return &CoinService{lgr: lgr, repo: repo}
}

func (srv *CoinService) SendCoins(senderID, receiverID, quantity int) error {
	srv.lgr.Info("Sending coins")

	if quantity <= 0 {
		srv.lgr.Error("Quantity is <= 0")
		return errors.New("quantity is <= 0")
	}

	tx, err := srv.repo.BeginTx()
	if err != nil {
		srv.lgr.Error("Error while beginning transaction", "error", err)
		return errors.New("error while beginning transaction")
	}
	defer tx.Rollback()

	fromUserID, err := srv.repo.GetUserForUpdate(senderID)
	if err != nil {
		srv.lgr.Error("Error while getting user for update", "error", err)
		return errors.New("error while getting user for update")
	}

	toUserID, err := srv.repo.GetUserForUpdate(receiverID)
	if err != nil {
		srv.lgr.Error("Error while getting user for update", "error", err)
		return errors.New("error while getting user for update")
	}

	if fromUserID.Coins < quantity {
		srv.lgr.Error("Not enough coins to send", "error", err)
		return errors.New("not enough coins to send")
	}

	if err = srv.repo.UpdateBalance(senderID, fromUserID.Coins-quantity); err != nil {
		srv.lgr.Error("Error while updating balance", "error", err)
		return errors.New("Error while updating balance")
	}

	if err = srv.repo.UpdateBalance(receiverID, toUserID.Coins+quantity); err != nil {
		srv.lgr.Error("Error while updating balance", "error", err)
		return errors.New("error while updating balance")
	}

	coinTransaction := &model.CoinTransaction{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Quantity:   quantity,
	}

	if err = srv.repo.CreateTransaction(coinTransaction); err != nil {
		srv.lgr.Error("Error while creating transaction", "error", err)
		return err
	}

	return tx.Commit()
}
