package service

import (
	"errors"
	"github.com/mdvasilyev/avito-shop/internal/model"
	"github.com/mdvasilyev/avito-shop/internal/repository"
	"log/slog"
)

type UserService struct {
	lgr  *slog.Logger
	repo *repository.UserRepository
}

func NewUserService(lgr *slog.Logger, repo *repository.UserRepository) *UserService {
	return &UserService{lgr: lgr, repo: repo}
}

func (srv *UserService) Info(userId int) (*model.UserResponse, error) {
	srv.lgr.Info("Getting info")

	coins, err := srv.repo.GetUserCoins(userId)
	if err != nil {
		srv.lgr.Error("Error while getting user coins", "error", err)
		return nil, errors.New("error while getting user coins")
	}

	inventory, err := srv.repo.GetUserInventory(userId)
	if err != nil {
		srv.lgr.Error("Error while getting user inventory", "error", err)
		return nil, errors.New("error while getting user inventory")
	}

	coinHistory, err := srv.repo.GetCoinTransactions(userId)
	if err != nil {
		srv.lgr.Error("Error while getting coin transactions", "error", err)
		return nil, errors.New("error while getting coin transactions")
	}

	return &model.UserResponse{
		Coins:       coins,
		Inventory:   inventory,
		CoinHistory: coinHistory,
	}, nil
}
