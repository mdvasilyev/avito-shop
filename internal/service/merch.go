package service

import (
	"errors"
	"github.com/mdvasilyev/avito-shop/internal/repository"
	"log/slog"
)

type MerchService struct {
	lgr  *slog.Logger
	repo *repository.MerchRepository
}

func NewMerchService(lgr *slog.Logger, repo *repository.MerchRepository) *MerchService {
	return &MerchService{lgr: lgr, repo: repo}
}

func (srv *MerchService) BuyItem(userID int, itemName string) error {
	srv.lgr.Info("Buying item")

	tx, err := srv.repo.BeginTx()
	if err != nil {
		srv.lgr.Error("Error while beginning transaction", "error", err)
		return errors.New("error while beginning transaction")
	}
	defer tx.Rollback()

	user, err := srv.repo.GetUserById(userID)
	if err != nil {
		srv.lgr.Error("Error while getting user by id", "error", err)
		return errors.New("user not found")
	}

	item, err := srv.repo.GetItemByName(itemName)
	if err != nil {
		srv.lgr.Error("Error while getting item by name", "error", err)
		return errors.New("item not found")
	}

	if user.Coins < item.Price {
		srv.lgr.Error("Not enough coins to buy")
		return errors.New("not enough coins to buy")
	}

	if err = srv.repo.UpdateUserBalance(userID, user.Coins-item.Price); err != nil {
		srv.lgr.Error("Error while updating user balance", "error", err)
		return errors.New("error while updating user balance")
	}

	if err = srv.repo.UpdateInventory(userID, item.ID); err != nil {
		srv.lgr.Error("Error while updating user inventory", "error", err)
		return errors.New("error while updating user inventory")
	}

	return tx.Commit()
}
